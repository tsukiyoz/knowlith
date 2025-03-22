package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsukiyoz/knowlith/internal/pkg/errorsx"
)

type ErrorResponse struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteResponse(c *fiber.Ctx, data any, err error) error {
	if err != nil {
		errx := errorsx.FromError(err)
		return c.Status(errx.Code).JSON(ErrorResponse{
			Reason:  errx.Reason,
			Message: errx.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
