package go_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
	assert := assert.New(t)
	cache := NewCache(5, LFU)

	assert.Equal(cache.Get("a"), nil)
	assert.Equal(cache.Get("b"), nil)
	assert.Equal(cache.Get("c"), nil)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	assert.Equal(cache.Get("a"), 1)
	assert.Equal(cache.Get("b"), 2)
	assert.Equal(cache.Get("c"), 3)
	cache.Set("d", 4)
	cache.Set("e", 5)
	assert.Equal(cache.Get("a"), 1)
	assert.Equal(cache.Get("b"), 2)
	assert.Equal(cache.Get("c"), 3)
	assert.Equal(cache.Get("d"), 4)
	assert.Equal(cache.Get("e"), 5)

	cache.Set(1, "abc")
	assert.Equal(cache.Get(1), "abc")
	assert.Equal(cache.Get("d"), nil)
	cache.Set(2, "pqr")
	assert.Equal(cache.Get(2), "pqr")
	assert.Equal(cache.Get("e"), nil)
	assert.Equal(cache.Get("a"), 1)
	assert.Equal(cache.Get("b"), 2)
	assert.Equal(cache.Get("c"), 3)

	cache.Get("b")
	cache.Get("b")
	cache.Get("b")
	cache.Get("b")
	cache.Get("b")
	cache.Set(3, "3")
	cache.Set(4, "4")
	cache.Set(5, "5")
	cache.Set(6, "6")
	cache.Set(7, "7")
	cache.Set(8, "8")

	assert.Equal(cache.Get("a"), 1)
	assert.Equal(cache.Get("b"), 2)
	assert.Equal(cache.Get("c"), 3)
	assert.Equal(cache.Get("d"), nil)
	assert.Equal(cache.Get("e"), nil)
	assert.Equal(cache.Get(1), nil)
	assert.Equal(cache.Get(2), "pqr")
	assert.Equal(cache.Get(3), nil)
	assert.Equal(cache.Get(4), nil)
	assert.Equal(cache.Get(5), nil)
	assert.Equal(cache.Get(6), nil)
	assert.Equal(cache.Get(7), nil)
	assert.Equal(cache.Get(8), "8")
}
