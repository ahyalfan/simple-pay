package main

import (
	"ahyalfan/golang_e_money/internal/api"
	"ahyalfan/golang_e_money/internal/component"
	"ahyalfan/golang_e_money/internal/config"
	"ahyalfan/golang_e_money/internal/middleware"
	"ahyalfan/golang_e_money/internal/repository"
	"ahyalfan/golang_e_money/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	app := fiber.New()

	// connection
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	// repository
	userRepository := repository.NewUser(dbConnection)

	// service
	emailService := service.NewEmail(cnf)
	userService := service.NewUserService(userRepository, cacheConnection, emailService)

	// middleware
	authMiddleware := middleware.Authenticate(userService)

	api.NewAuth(app, userService, authMiddleware)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		log.Fatal(err.Error())
	}

}
