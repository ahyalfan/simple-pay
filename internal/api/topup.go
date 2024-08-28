package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type topupApi struct {
	topupService domain.TopupService
}

func NewTopup(app *fiber.App, authMid fiber.Handler, topupService domain.TopupService) {
	api := topupApi{topupService: topupService}

	app.Post("/api/topup/initialize", authMid, api.initializeTopup)
}

func (ta *topupApi) initializeTopup(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TopupReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}
	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(fiber.StatusBadRequest, "validates failed", fails))
	}

	user := ctx.Locals("x-user").(dto.UserData)
	req.UserId = user.ID
	res, err := ta.topupService.InitializeTopup(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateError(fiber.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(dto.CreateSuccess(fiber.StatusOK, "topup request initiated", res))
}
