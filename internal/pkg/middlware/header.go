package middlware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors fiber.Handler = cors.New(cors.Config{
	AllowOrigins: "*",
	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
	}, ","),
	Next: func(c *fiber.Ctx) bool {
		return c.Method() == "OPTIONS"
	},
})

var NoCache fiber.Handler = cache.New(cache.Config{
	Next: func(c *fiber.Ctx) bool {
		return c.Query("noCache") == "true"
	},
	Expiration:   30 * time.Minute,
	CacheControl: true,
})
