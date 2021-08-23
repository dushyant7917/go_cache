package go_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIFOCache(t *testing.T) {
	assert := assert.New(t)
	cache := NewCache(3, FIFO)

	assert.Equal(cache.Get("dushyant"), nil)
	cache.Set("dushyant", 25)
	assert.Equal(cache.Get("dushyant"), 25)

	assert.Equal(cache.Get(42), nil)
	cache.Set(42, 68.47)
	assert.Equal(cache.Get(42), 68.47)

	cache.Set(42, "xyz")
	assert.Equal(cache.Get(42), "xyz")

	cache.Set("foo", "bar")
	assert.Equal(cache.Get("foo"), "bar")

	cache.Set(7, 67)
	assert.Equal(cache.Get(7), 67)
	assert.Equal(cache.Get("dushyant"), nil)

	cache.Set(9, "foo")
	assert.Equal(cache.Get(9), "foo")
	assert.Equal(cache.Get(42), nil)
}
