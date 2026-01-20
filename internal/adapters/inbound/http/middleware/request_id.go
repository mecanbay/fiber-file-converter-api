package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-Id"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {

		reqId := c.Get(RequestIDHeader)
		if reqId == "" {
			reqId = uuid.NewString()
		}
		c.Locals("request_id", reqId)
		c.Set(RequestIDHeader, reqId)

		return c.Next()
	}
}
