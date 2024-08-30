package api

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	userService domain.UserService
	fdsService  domain.FdsService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMid fiber.Handler, fdsService domain.FdsService) {
	handler := authApi{userService: userService, fdsService: fdsService}

	app.Post("/api/auth", handler.GenerateToken)
	app.Get("/api/auth/validate", authMid, handler.ValidateToken)
	app.Post("/api/auth/register", handler.Register)
	app.Post("/api/auth/validate-otp", handler.ValidateOTP)
}

func (a *authApi) GenerateToken(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}

	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(400, "validates failed", fails))
	}

	token, err := a.userService.Authenticate(c, req)
	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.CreateError(401, err.Error()))
	}

	// kita bisa ambil ip addressnya
	// karena ini host nya masih local host kita belum bisa implementasikan yg real
	// disini kita coba pakai permisalan ip di Header
	if !a.fdsService.IsAuthorized(c, ctx.Get("X-FORWARDED-FOR"), token.UserID) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(401, "unauthorized"))
	}
	return ctx.JSON(dto.CreateSuccess(200, "success", token))
}

func (a *authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.JSON(dto.CreateSuccess(400, "valid token", user))
}

func (a *authApi) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UserRegisterReg
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}
	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(400, "validates failed", fails))
	}
	res, err := a.userService.Register(c, req)
	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.CreateError(400, err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.CreateSuccess(fiber.StatusCreated, "success created", res))
}

func (a *authApi) ValidateOTP(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.ValidateOtpReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(400, err.Error()))
	}
	fails := util.Vallidate(req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateResponseErrorData(400, "validates failed", fails))
	}
	err := a.userService.ValidateOTP(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(dto.CreateError(400, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, "token Valid", ""))
}
