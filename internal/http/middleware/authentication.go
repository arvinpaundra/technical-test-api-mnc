package middleware

import (
	"strings"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/app/auth"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/format"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(f *factory.Factory) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")

		token := strings.ReplaceAll(bearerToken, "Bearer ", "")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(format.Failed("Unauthenticated"))
		}

		claims, err := util.DecodeJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(format.Failed("Unauthenticated"))
		}

		user, err := auth.NewService(f).Authenticate(c.Context(), claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(format.Failed("Unauthenticated"))
		}

		c.Locals("claims", user)

		return c.Next()
	}
}
