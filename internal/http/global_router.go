package http

import (
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/http/auth"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/http/middleware"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/http/transaction"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/http/user"
)

func NewHttpRouter(app *fiber.App, f *factory.Factory) {
	// setup app middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-type, Authorization",
	}))

	app.Use(recover.New())

	// index route
	app.Get("/", index)

	auth.NewHandler(f).RouterV1(app.Group("/"))

	// guarded api routes
	guard := app.Group("", middleware.Authenticate(f))

	transaction.NewHandler(f).RouterV1(guard)
	user.NewHandler(f).RouterV1(guard)
}

func index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"name":    "technical-test-api-mnc",
		"version": util.LoadVersion(),
	})
}
