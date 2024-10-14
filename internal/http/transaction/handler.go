package transaction

import (
	"context"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/app/transaction"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/format"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Topup(ctx context.Context, userId string, payload request.Topup) (response.Topup, error)
	Payment(ctx context.Context, userId string, payload request.Payment) (response.Payment, error)
	Transfer(ctx context.Context, userId string, payload request.Transfer) (response.Transfer, error)
	GetTransactions(ctx context.Context, userId string) ([]response.Transaction, error)
}

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: transaction.NewService(f),
	}
}

func (h *Handler) HandlerTopup(c *fiber.Ctx) error {
	var payload request.Topup

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	claims := c.Locals("claims").(response.Authenticate)

	res, err := h.service.Topup(c.Context(), claims.UserId, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(format.Success(res))
}

func (h *Handler) HandlerPayment(c *fiber.Ctx) error {
	var payload request.Payment

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	claims := c.Locals("claims").(response.Authenticate)

	res, err := h.service.Payment(c.Context(), claims.UserId, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(format.Success(res))
}

func (h *Handler) HandlerTransfer(c *fiber.Ctx) error {
	var payload request.Transfer

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	claims := c.Locals("claims").(response.Authenticate)

	res, err := h.service.Transfer(c.Context(), claims.UserId, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(format.Success(res))
}

func (h *Handler) HandlerGetTransactions(c *fiber.Ctx) error {
	claims := c.Locals("claims").(response.Authenticate)

	res, err := h.service.GetTransactions(c.Context(), claims.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(format.Success(res))
}
