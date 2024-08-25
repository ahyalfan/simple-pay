package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type notificationApi struct {
	notificationService domain.NotificationService
}

func NewNotification(app *fiber.App, notificationService domain.NotificationService, authMid fiber.Handler) {
	api := notificationApi{notificationService: notificationService}

	app.Get("/api/notifications", authMid, api.getNotificationByUser)
}

func (api *notificationApi) getNotificationByUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("x-user").(dto.UserData)
	notifications, err := api.notificationService.FindByUser(c, user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateError(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(dto.CreateSuccess(fiber.StatusOK, "notifications fetched", notifications))
}
