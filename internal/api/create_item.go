package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

type CreateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) CreateItem(c *fiber.Ctx) error {
	var req CreateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	err := s.Repository.Create(c.Context(), trackerstore.Item{
		ID:   uuid.New().String(),
		Name: req.Name,
	})
	if err != nil {
		log.Error(trackerstore.ErrCreate(err))
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusCreated)
}
