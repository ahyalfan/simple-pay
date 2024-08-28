package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topupService    domain.TopupService
}

func NewMidtrans(app *fiber.App, midtransService domain.MidtransService, topupService domain.TopupService) {
	api := midtransApi{midtransService: midtransService, topupService: topupService}
	app.Post("/api/midtrans/payment-callback", api.paymentHandlerNotification)
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(dto.CreateSuccess(200, "test success", "test message"))
	})
}

func (api *midtransApi) paymentHandlerNotification(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var notificationPayload map[string]interface{}
	if err := ctx.BodyParser(&notificationPayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, err.Error()))
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, "missing key `order_id"))
	}

	result, err := api.midtransService.VerifyPayment(c, orderId)
	if result {
		err = api.topupService.ConfimedTopup(c, orderId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateError(fiber.StatusInternalServerError, err.Error()))
		}
		return ctx.SendStatus(200)
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateError(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.SendStatus(400)
}
