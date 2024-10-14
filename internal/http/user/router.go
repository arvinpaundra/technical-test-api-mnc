package user

import "github.com/gofiber/fiber/v2"

func (h *Handler) RouterV1(app fiber.Router) {
	app.Put("/profile", h.HandlerUpdateProfile)
}
