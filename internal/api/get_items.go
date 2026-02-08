package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

type ItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetItemsResponse struct {
	Items []ItemRequest `json:"items"`
}

func (s *Server) GetItems(c *fiber.Ctx) error {
	if c.Query("find", "") != "" {
		return s.FindItem(c)
	}
	items, err := s.Repository.List(c.Context())
	if err != nil {
		log.Error(trackerstore.ErrGet(err))
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{Items: res})
}
