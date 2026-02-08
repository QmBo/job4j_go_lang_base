package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id required")
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	item, err := s.Repository.Get(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("item with id %s not found", id))
	}
	if err := s.Repository.UpdateByID(c.Context(), trackerstore.Item{ID: item.ID, Name: req.Name}); err != nil {
		log.Error(trackerstore.ErrUpdate(err))
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return c.SendStatus(fiber.StatusOK)
}
