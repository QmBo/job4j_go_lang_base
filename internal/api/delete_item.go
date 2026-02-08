package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id required")
	}
	if err := s.Repository.DeleteByID(c.Context(), id); err != nil {
		log.Error(trackerstore.ErrDelete(err))
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
