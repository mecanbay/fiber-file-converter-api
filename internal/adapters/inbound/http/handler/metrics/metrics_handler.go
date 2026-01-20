package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) SetupRoutes(app *fiber.App) {
	app.Get("/metrics", adaptor.HTTPHandler(h.Metrics()))
}
