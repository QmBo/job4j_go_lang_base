package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].ID = "2"

		assert.Equal(t,
			"1",
			tracker.GetItems()[0].ID,
		)
	})
	t.Run("get index", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		res := tracker.GetItems()
		assert.Equal(t, 1, len(res))
		indexOf := tracker.IndexOf(item)
		assert.Equal(t, 0, indexOf)
	})
	t.Run("get item", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		item = base.Item{
			ID:   "2",
			Name: "Second Item",
		}
		tracker.AddItem(item)
		res := tracker.GetItems()
		assert.Equal(t, 2, len(res))
		assert.Equal(t, item, res[1])
		assert.NotEqual(t, item, res[0])
	})
	t.Run("delete item", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		item = base.Item{
			ID:   "2",
			Name: "Second Item",
		}
		tracker.AddItem(item)
		item2 := base.Item{
			ID:   "3",
			Name: "Third Item",
		}
		tracker.AddItem(item2)

		res := tracker.IndexOf(item)
		assert.Equal(t, 1, res)

		removeItem := tracker.RemoveItem(res)
		assert.Equal(t, item, removeItem)

		items := tracker.GetItems()
		assert.Equal(t, 2, len(items))
		assert.Equal(t, items, []base.Item{
			{
				ID:   "1",
				Name: "First Item",
			},
			{
				ID:   "3",
				Name: "Third Item",
			},
		})
	})
	t.Run("update item", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)
		item2 := base.Item{
			ID:   "2",
			Name: "Second Item",
		}
		tracker.Update(0, item2)

		assert.Equal(t, item2, tracker.GetItems()[0])
	})
}
