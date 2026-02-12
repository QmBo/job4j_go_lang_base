package base

import "fmt"

type Item struct {
	ID   string
	Name string
}
type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) GetItem(index int) Item {
	return t.items[index]
}

func (t *Tracker) IndexOf(item Item) (int, error) {
	for i := range t.items {
		if t.items[i].ID == item.ID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("item not found")
}

func (t *Tracker) Update(item Item) ([]Item, error) {
	i, err := t.IndexOf(item)
	if err != nil {
		return []Item{}, fmt.Errorf("item not found")
	}
	t.items[i] = item
	return t.GetItems(), nil
}

func (t *Tracker) RemoveItem(index int) Item {
	res := t.items[index]
	t.items = append(t.items[:index], t.items[index+1:]...)
	return res
}
