package health

import (
	app "fiber-file-converter-api/internal/application/health"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *app.Service
}

func NewHandler(service *app.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRoutes(app *fiber.App) {
	app.Get("/health", h.Health)
}
