package go_cache

import (
	"container/list"
)

type KeyVal struct {
	key interface{}
	val interface{}
}

type CapacityDetails struct {
	maxKeyCount int
}

type Cache interface {
	Set(interface{}, interface{})
	Get(interface{}) interface{}
	Evict()
	GetCurrentKeyCount() int
	Full() bool
}

func NewCache(maxKeyCount int, cacheType CacheType) Cache {
	if cacheType == FIFO {
		return &FIFOCache{
			CapacityDetails: CapacityDetails{maxKeyCount: maxKeyCount},
			KeyValPairs:     list.New(),
			KeyToKeyVal:     make(map[interface{}](*list.Element)),
		}
	}
	if cacheType == LRU {
		return &LRUCache{
			CapacityDetails: CapacityDetails{maxKeyCount: maxKeyCount},
			KeyValPairs:     list.New(),
			KeyToKeyVal:     make(map[interface{}](*list.Element)),
		}
	}
	return nil
}
