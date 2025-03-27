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
		e := errorsx.FromError(err)
		return c.Status(e.Code).JSON(ErrorResponse{
			Reason:  e.Reason,
			Message: e.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
