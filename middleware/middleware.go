package middleware

import (
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "UNAUTHENTICATED",
		})
	}

	_, err := utils.VerifyTokenJwt(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "UNAUTHENTICATED",
		})
	}
	
	return ctx.Next()
}

func IsAdmin(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "UNAUTHENTICATED",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message" : "FORBIDDEN ACCESS",
		})
	}

	// ctx.Locals("userInfo", claims)

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}