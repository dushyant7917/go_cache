package go_cache

import (
	"testing"
)

func TestNewCache(t *testing.T) {
	_ = NewCache(3, FIFO)
	_ = NewCache(5, LRU)
}
