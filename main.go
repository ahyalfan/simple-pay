package main

import (
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/api"
	"ahyalfan/golang_e_money/internal/component"
	"ahyalfan/golang_e_money/internal/config"
	"ahyalfan/golang_e_money/internal/middleware"
	"ahyalfan/golang_e_money/internal/repository"
	"ahyalfan/golang_e_money/internal/service"
	"ahyalfan/golang_e_money/internal/sse"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	app := fiber.New()

	// connection
	dbConnection := component.GetDatabaseConnection(cnf)
	// cacheConnection := component.GetCacheConnection() // kita coba ganti dengan redis
	cacheConnection := repository.NewRedisCache(cnf)

	// hub
	hub := &dto.Hub{
		NotificationChannel: make(map[int64]chan dto.NotificationData),
	}

	// repository
	userRepository := repository.NewUser(dbConnection)
	accountRepository := repository.NewAccount(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)
	notificationRepository := repository.NewNotification(dbConnection)
	templateRepository := repository.NewTemplate(dbConnection)
	topupRepository := repository.NewTopup(dbConnection)
	factorRepository := repository.NewFactor(dbConnection)

	// service
	emailService := service.NewEmail(cnf)
	factorService := service.NewFactor(factorRepository)
	userService := service.NewUserService(userRepository, cacheConnection, emailService, factorService)
	accountService := service.NewAccount(accountRepository)
	notificationService := service.NewNotification(notificationRepository, templateRepository, hub)
	transactionService := service.NewTransaction(accountRepository, transactionRepository, cacheConnection, dbConnection, notificationService)
	midtransService := service.NewMidtrans(cnf)
	topupService := service.NewTopupService(notificationService, topupRepository, midtransService, accountRepository, transactionRepository)

	// middleware
	authMiddleware := middleware.Authenticate(userService)

	api.NewAuth(app, userService, authMiddleware)
	api.NewMidtrans(app, midtransService, topupService)
	api.NewAccount(app, accountService, authMiddleware)
	api.NewTransfer(app, transactionService, authMiddleware, factorService)
	api.NewNotification(app, notificationService, authMiddleware)
	api.NewTopup(app, authMiddleware, topupService)

	// sse
	sse.NewNotificationSse(app, hub, authMiddleware)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		log.Fatal(err.Error())
	}

}
