package handler

import (
	"be-service-public-api/domain"

	"github.com/gofiber/fiber/v2"
)

type PublicHandler struct {
	PublicAPIUseCase   domain.PublicAPIUseCase
	PublicAPIMySQLRepo domain.PublicAPIMySQLRepo
}

func (ph *PublicHandler) HandlerFunction(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
