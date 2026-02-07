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
}
