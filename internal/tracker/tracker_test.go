package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("Add same id - error", func(t *testing.T) {
		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}

		addItem, err := tracker.AddItem(item)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, addItem.ID)

		item = Item{
			ID:   "1",
			Name: "David Doe",
		}
		_, err = tracker.AddItem(item)
		assert.Equal(t, ErrIDAlreadyExists, err)

		item = Item{
			ID:   "2",
			Name: "John Doe",
		}
		_, err = tracker.AddItem(item)
		assert.Nil(t, err)
	})
	t.Run("Add get - ok", func(t *testing.T) {
		t.Parallel()
		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		addItem, err := tracker.AddItem(item)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, addItem.ID)

		items := tracker.GetItems()

		assert.Equal(t, 1, len(items))
		assert.Equal(t, item.ID, items[0].ID)
	})
	t.Run("Add find - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		addItem, err := tracker.AddItem(item)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, addItem.ID)

		find, err := tracker.Find("name")
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, []Item{}, find)
	})
	t.Run("Add find - found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()

		items, err := tracker.Find("name")
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, []Item{}, items)

		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		addItem, err := tracker.AddItem(item)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, addItem.ID)

		find, err := tracker.Find("Doe")
		assert.Nil(t, err)
		assert.NotNil(t, find)
		assert.Equal(t, item.ID, find[0].ID)
		assert.Equal(t, item.Name, find[0].Name)
	})
	t.Run("Add get by position - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)

		find, err := tracker.GetByPosition(2)
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, Item{}, find)
	})
	t.Run("Add get by position - found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)

		find, err := tracker.GetByPosition(1)
		assert.Nil(t, err)
		assert.NotNil(t, find)
		assert.Equal(t, item.ID, find.ID)
		assert.Equal(t, item.Name, find.Name)
	})
	t.Run("Remove - success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)
		assert.Nil(t, err)
		removeItem, err := tracker.RemoveItem(1)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, removeItem.ID)
		items := tracker.GetItems()
		assert.Equal(t, 0, len(items))
	})
	t.Run("Remove - error", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)
		assert.Nil(t, err)
		item2 := Item{
			ID:   "2",
			Name: "John Doe",
		}
		_, err = tracker.AddItem(item2)
		assert.Nil(t, err)
		removeItem, err := tracker.RemoveItem(99)
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, Item{}, removeItem)
		removeItem, err = tracker.RemoveItem(1)
		assert.Nil(t, err)
		assert.Equal(t, item.ID, removeItem.ID)

		position, err := tracker.GetByPosition(1)
		assert.Nil(t, err)
		assert.Equal(t, item2.ID, position.ID)
	})
	t.Run("update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)
		assert.Nil(t, err)

		item2 := Item{
			ID:   "2",
			Name: "Doe John",
		}
		update, err := tracker.Update(item2.ID, item2)
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, Item{}, update)
	})
	t.Run("update - success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)
		assert.Nil(t, err)

		item2 := Item{
			ID:   "2",
			Name: "Doe John",
		}
		update, err := tracker.Update(item.ID, item2)
		assert.Nil(t, err)
		assert.Equal(t, item2.ID, update.ID)
		assert.Equal(t, item2.Name, update.Name)
		items := tracker.GetItems()
		assert.Equal(t, 1, len(items))
	})
	t.Run("String - success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "John Doe",
		}
		_, err := tracker.AddItem(item)
		assert.Nil(t, err)
		getItem, err := tracker.GetByPosition(1)
		assert.Nil(t, err)
		assert.Equal(t, "1) 1\tJohn Doe", getItem.toString())
	})
}
