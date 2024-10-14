package user

import (
	"context"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/app/user"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/format"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	UpdateProfile(ctx context.Context, userId string, payload request.UpdateProfile) (response.UpdateProfile, error)
}

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: user.NewService(f),
	}
}

func (h *Handler) HandlerUpdateProfile(c *fiber.Ctx) error {
	var payload request.UpdateProfile

	_ = c.BodyParser(&payload)

	validationError := validator.Validate(payload, validator.JSON)
	if validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(format.BadRequest(validationError))
	}

	claims := c.Locals("claims").(response.Authenticate)

	res, err := h.service.UpdateProfile(c.Context(), claims.UserId, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(format.Failed(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(format.Success(res))
}
