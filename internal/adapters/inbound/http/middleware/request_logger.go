package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ignore browser noise
		path := c.Path()
		if path == "/favicon.ico" || strings.HasPrefix(path, "/.well-known/") {
			return c.Next()
		}

		start := time.Now()

		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()
		requestId, _ := c.Locals("request_id").(string)

		fields := []zap.Field{
			zap.String("request_id", requestId),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
			zap.String("user_agent", string(c.Request().Header.UserAgent())),
		}

		switch {
		case status >= 500:
			zap.L().Error("http request", fields...)
		case status >= 400:
			zap.L().Warn("http request", fields...)
		default:
			zap.L().Info("http request", fields...)
		}

		return err
	}
}
