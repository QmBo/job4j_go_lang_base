package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

func (s *Server) FindItem(c *fiber.Ctx) error {
	params := c.Query("find", "")
	if params == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param is required")
	}
	find, err := s.Repository.Find(c.Context(), params)
	if err != nil {
		log.Error(trackerstore.ErrFind(err))
		return fiber.NewError(fiber.StatusInternalServerError, "cannot find item")
	}
	res := make([]ItemRequest, 0, len(find))
	for _, item := range find {
		ir := ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		}
		res = append(res, ir)
	}
	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{res})
}
