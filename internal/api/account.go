package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type accountApi struct {
	accountService domain.AccountService
}

func NewAccount(app *fiber.App, accountService domain.AccountService, authMid fiber.Handler) {
	api := accountApi{
		accountService: accountService,
	}

	app.Post("/api/account/create", authMid, api.Create)
}

func (a *accountApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := a.accountService.Create(c)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(dto.CreateError(fiber.StatusBadGateway, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, "account created", res))

}
