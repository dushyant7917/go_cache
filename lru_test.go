package go_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	assert := assert.New(t)
	cache := NewCache(2, LRU)

	assert.Equal(cache.Get("dushyant"), nil)
	assert.Equal(cache.GetCurrentKeyCount(), 0)
	cache.Set("dushyant", 25)
	assert.Equal(cache.GetCurrentKeyCount(), 1)
	assert.Equal(cache.Get("dushyant"), 25)

	assert.Equal(cache.Get(42), nil)
	cache.Set(42, 68.47)
	assert.Equal(cache.GetCurrentKeyCount(), 2)
	assert.Equal(cache.Get(42), 68.47)

	cache.Get("dushyant")
	cache.Set("foo", "bar")
	assert.Equal(cache.Get("foo"), "bar")
	assert.Equal(cache.Get("dushyant"), 25)
	assert.Equal(cache.Get(42), nil)
	assert.Equal(cache.GetCurrentKeyCount(), 2)
}
