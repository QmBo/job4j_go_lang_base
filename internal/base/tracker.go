package base

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

func (t *Tracker) IndexOf(item Item) int {
	for i := range t.items {
		if t.items[i] == item {
			return i
		}
	}
	return -1
}

func (t *Tracker) Update(item Item) []Item {
	for i := range t.items {
		if t.items[i].ID == item.ID {
			t.items[i] = item
		}
	}
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) RemoveItem(index int) Item {
	res := t.items[index]
	t.items = append(t.items[:index], t.items[index+1:]...)
	return res
}
