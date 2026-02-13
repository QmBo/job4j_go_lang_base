package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID       string
	Name     string
	position int
}

func (i Item) toString() string {
	return fmt.Sprintf("%d) %s\t%s", i.position, i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, err := t.findById(item.ID)
	if err == nil {
		return Item{}, ErrIDAlreadyExists
	}
	item.position = len(t.items) + 1
	t.items = append(t.items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) GetByPosition(position int) (Item, error) {
	for _, item := range t.items {
		if item.position == position {
			return item, nil
		}
	}
	return Item{}, ErrNotFound
}

func (t *Tracker) Update(id string, item Item) (Item, error) {
	for i := range t.items {
		if t.items[i].ID == id {
			t.items[i].position = item.position
			return item, nil
		}
	}
	return Item{}, ErrNotFound
}

func (t *Tracker) RemoveItem(position int) (Item, error) {
	res, err := t.GetByPosition(position)
	if err != nil {
		return Item{}, err
	}
	index := res.position - 1
	t.items = append(t.items[:index], t.items[index+1:]...)
	for i := range t.items {
		t.items[i].position = i + 1
	}
	return res, nil
}

func (t *Tracker) Find(name string) ([]Item, error) {
	if len(t.items) == 0 {
		return []Item{}, ErrNotFound
	}
	res := make([]Item, 0, len(t.items))
	for _, item := range t.items {
		if strings.Contains(item.Name, name) {
			res = append(res, item)
		}
	}
	if len(res) == 0 {
		return []Item{}, ErrNotFound
	}
	return res, nil
}

func (t *Tracker) findById(id string) (Item, error) {
	for _, item := range t.items {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, ErrNotFound
}
