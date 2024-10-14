package auth

import "github.com/gofiber/fiber/v2"

func (h *Handler) RouterV1(app fiber.Router) {
	app.Post("/register", h.HandlerRegister)
	app.Post("/login", h.HandlerLogin)
}
