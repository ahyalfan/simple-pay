package middleware

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// token di taruh di bearer token
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.CreateError(fiber.StatusUnauthorized, "Token required"))
		}

		user, err := userService.ValidateToken(ctx.Context(), token)
		if err != nil {
			return ctx.Status(util.GetHttpStatus(err)).JSON(dto.CreateError(fiber.StatusUnauthorized, "Token invalid"))
		}
		ctx.Locals("x-user", user) // simpan di memory context

		return ctx.Next()
	}
}
