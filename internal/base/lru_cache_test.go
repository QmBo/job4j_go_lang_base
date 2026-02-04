package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_LruCache(t *testing.T) {
	t.Parallel()

	t.Run("put anf get", func(t *testing.T) {
		t.Parallel()
		cache, _ := base.NewLruCache(1)

		cache.Put("hello", "world")

		expected := "world"
		assert.Equal(t, &expected, cache.Get("hello"))
		assert.Nil(t, cache.Get("world"))
		assert.Equal(t, expected, cache.Head.Value)
		assert.Equal(t, expected, cache.Tail.Value)
	})
	t.Run("add 2 for capacity 1", func(t *testing.T) {
		t.Parallel()
		cache, _ := base.NewLruCache(1)

		cache.Put("1", "3")
		cache.Put("2", "4")

		rsl := cache.Get("2")
		expected := "4"
		assert.Equal(t, &expected, rsl)
		assert.Nil(t, cache.Get("1"))
	})
	t.Run("add 2 get - 2", func(t *testing.T) {
		t.Parallel()
		cache, _ := base.NewLruCache(2)

		cache.Put("1", "3")
		cache.Put("2", "4")

		expected := "3"
		headExpKey := "2"
		tailExpKey := "1"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
		assert.Equal(t, &expected, cache.Get("1"))
		headExpKey = "1"
		tailExpKey = "2"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
		expected = "4"
		assert.Equal(t, &expected, cache.Get("2"))
		headExpKey = "2"
		tailExpKey = "1"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
		assert.Nil(t, cache.Get("5"))
	})
	t.Run("add 3 and revers", func(t *testing.T) {
		t.Parallel()
		cache, _ := base.NewLruCache(5)

		cache.Put("3", "c")
		cache.Put("2", "b")
		cache.Put("1", "a")

		expected := "a"
		headExpKey := "1"
		tailExpKey := "3"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
		assert.Equal(t, &expected, cache.Get("1"))
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)

		expected = "b"
		headExpKey = "2"
		tailExpKey = "3"
		assert.Equal(t, &expected, cache.Get("2"))
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)

		expected = "c"
		headExpKey = "3"
		tailExpKey = "1"
		assert.Equal(t, &expected, cache.Get("3"))
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
	})
	t.Run("over capacity", func(t *testing.T) {
		t.Parallel()
		cache, _ := base.NewLruCache(3)

		cache.Put("4", "e")
		cache.Put("3", "c")
		cache.Put("2", "b")
		cache.Put("1", "a")

		expected := "c"
		headExpKey := "1"
		tailExpKey := "3"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
		assert.Equal(t, &expected, cache.Get("3"))

		headExpKey = "3"
		tailExpKey = "2"
		assert.Equal(t, headExpKey, cache.Head.Key)
		assert.Equal(t, tailExpKey, cache.Tail.Key)
	})
	t.Run("incorrect capacity", func(t *testing.T) {
		t.Parallel()
		_, err := base.NewLruCache(0)
		assert.NotNil(t, err)
		assert.Equal(t, "capacity must be greater than zero", err.Error())
	})
}
