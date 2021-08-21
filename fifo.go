package go_cache

import (
	"container/list"
)

type FIFOCache struct {
	CapacityDetails
	KeyValPairs *list.List
	KeyToKeyVal map[interface{}](*list.Element)
}

func (cache *FIFOCache) Set(key, val interface{}) {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		keyVal := kv.Value.(*KeyVal)
		keyVal.val = val
		return
	}

	if cache.Full() {
		cache.Evict()
	}
	cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
	cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
}

func (cache *FIFOCache) Get(key interface{}) interface{} {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		return kv.Value.(*KeyVal).val
	}

	return nil
}

func (cache *FIFOCache) Evict() {
	if len(cache.KeyToKeyVal) == 0 {
		return
	}

	front_kv := cache.KeyValPairs.Front()
	key := front_kv.Value.(*KeyVal).key
	cache.KeyValPairs.Remove(front_kv)
	delete(cache.KeyToKeyVal, key)
}

func (cache *FIFOCache) Full() bool {
	return len(cache.KeyToKeyVal) == cache.maxKeyCount
}

func (cache *FIFOCache) GetCurrentKeyCount() int {
	return len(cache.KeyToKeyVal)
}
