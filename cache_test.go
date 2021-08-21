package go_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	assert := assert.New(t)
	cache := NewCache(3, FIFO)
	_, ok := cache.(Cache)
	assert.Equal(ok, true)
}
