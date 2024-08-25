package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type transferApi struct {
	transactionService domain.TransactionService
}

func NewTransfer(app *fiber.App, transaction domain.TransactionService, authMid fiber.Handler) {
	handler := transferApi{transactionService: transaction}
	api := app.Group("/api", authMid)

	api.Post("/transfer/inquiry", handler.transferInquiry)
	api.Post("/transfer/execute", handler.transferExecute)

}

func (ta *transferApi) transferInquiry(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TransferInQuiryReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(fiber.StatusBadRequest, "validates failed", fails))
	}

	val, err := ta.transactionService.TransferInquiry(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(dto.CreateError(fiber.StatusBadGateway, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, "inquiry success", val))

}
func (ta *transferApi) transferExecute(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TransferExecuteReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(fiber.StatusBadRequest, "validates failed", fails))
	}

	err := ta.transactionService.TransferExecute(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(dto.CreateError(fiber.StatusBadGateway, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, "inquiry success", ""))
}
