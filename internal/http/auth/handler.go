package auth

import (
	"context"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/app/auth"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/format"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Register(ctx context.Context, payload request.Register) (response.Register, error)
	Login(ctx context.Context, payload request.Login) (response.Login, error)
	RenewRefreshToken(ctx context.Context) error
}

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: auth.NewService(f),
	}
}

func (h *Handler) HandlerRegister(c *fiber.Ctx) error {
	var payload request.Register

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	res, err := h.service.Register(c.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrPhoneAlreadyExist:
			return c.Status(fiber.StatusConflict).JSON(format.Failed(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(format.Success(res))
}

func (h *Handler) HandlerLogin(c *fiber.Ctx) error {
	var payload request.Login

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	res, err := h.service.Login(c.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrUserNotFound:
			return c.Status(fiber.StatusConflict).JSON(format.Failed(constant.ErrPhoneAndPinNotMatch.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(format.Success(res))
}
