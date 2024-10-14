package transaction

import "github.com/gofiber/fiber/v2"

func (h *Handler) RouterV1(app fiber.Router) {
	app.Post("/topup", h.HandlerTopup)
	app.Post("/pay", h.HandlerPayment)
	app.Post("/transfer", h.HandlerTransfer)
	app.Get("/transactions", h.HandlerGetTransactions)
}
