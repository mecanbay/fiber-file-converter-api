package health

import "github.com/gofiber/fiber/v2"

type GetHealtCheckResponse struct {
	Status string `json:"status"`
}

func (h Handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(GetHealtCheckResponse{Status: string(h.service.Check())})
}
