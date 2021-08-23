package go_cache

import (
	"container/list"
)

type KeyVal struct {
	key, val AnyType
}

type CapacityDetails struct {
	maxKeyCount int
}

type Cache interface {
	Set(AnyType, AnyType)
	Get(AnyType) AnyType
	Evict()
}

func NewCache(maxKeyCount int, cacheType CacheType) Cache {
	capacityDetails := CapacityDetails{maxKeyCount: maxKeyCount}
	dll := list.New()
	keyToItem := make(KeyToDLLElement)

	if cacheType == FIFO {
		return &FIFOCache{
			CapacityDetails: capacityDetails,
			KeyValPairs:     dll,
			KeyToKeyVal:     keyToItem,
		}
	}
	if cacheType == LRU {
		return &LRUCache{
			CapacityDetails: capacityDetails,
			KeyValPairs:     dll,
			KeyToKeyVal:     keyToItem,
		}
	}
	if cacheType == LFU {
		return &LFUCache{
			CapacityDetails: capacityDetails,
			Frequencies:     dll,
			KeyToKeyValFreq: keyToItem,
		}
	}
	return nil
}
