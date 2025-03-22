package middlware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/tsukiyoz/knowlith/internal/pkg/known"
)

var RequestID fiber.Handler = requestid.New(requestid.Config{
	Header: known.XRequestID,
	Generator: func() string {
		return uuid.New().String()
	},
})
