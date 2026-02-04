package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("valid - 0", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			UserID:      "id",
			Title:       "title",
			Description: "description",
		}
		rsl := base.Validate(&in)

		assert.Equal(t, 0, len(rsl))
	})
	t.Run("nil - 1 (nil is passed)", func(t *testing.T) {
		t.Parallel()

		var in *base.ValidateRequest

		rsl := base.Validate(in)

		assert.Equal(t, 1, len(rsl))
		assert.Equal(t, "nil is passed", rsl[0])
	})
	t.Run("all fields are empty - 3 (UserID is empty, Title is empty, Description is empty)", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{}
		rsl := base.Validate(&in)

		assert.Equal(t, 3, len(rsl))
		assert.Equal(t, "UserID is empty", rsl[0])
		assert.Equal(t, "Title is empty", rsl[1])
		assert.Equal(t, "Description is empty", rsl[2])
	})
	t.Run("two fields are empty - 2 (Title is empty, Description is empty)", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			UserID: "id",
		}
		rsl := base.Validate(&in)

		assert.Equal(t, 2, len(rsl))
		assert.Equal(t, "Title is empty", rsl[0])
		assert.Equal(t, "Description is empty", rsl[1])
	})
	t.Run("one field is empty - 1 (Description is empty)", func(t *testing.T) {
		t.Parallel()

		in := base.ValidateRequest{
			UserID: "1",
			Title:  "title",
		}
		rsl := base.Validate(&in)

		assert.Equal(t, 1, len(rsl))
		assert.Equal(t, "Description is empty", rsl[0])
	})

}
