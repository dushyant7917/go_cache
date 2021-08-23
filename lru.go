package go_cache

import (
	"container/list"
)

type LRUCache struct {
	CapacityDetails
	KeyValPairs *list.List
	KeyToKeyVal map[interface{}](*list.Element)
}

func (cache *LRUCache) Set(key, val interface{}) {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		cache.KeyValPairs.Remove(kv)
		cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
		cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
		return
	}

	if cache.Full() {
		cache.Evict()
	}
	cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
	cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
}

func (cache *LRUCache) Get(key interface{}) interface{} {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		val := kv.Value.(*KeyVal).val
		cache.KeyValPairs.Remove(kv)
		cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
		cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
		return val
	}

	return nil
}

func (cache *LRUCache) Evict() {
	if len(cache.KeyToKeyVal) == 0 {
		return
	}

	lru_kv := cache.KeyValPairs.Front()
	key := lru_kv.Value.(*KeyVal).key
	cache.KeyValPairs.Remove(lru_kv)
	delete(cache.KeyToKeyVal, key)
}

func (cache *LRUCache) Full() bool {
	return len(cache.KeyToKeyVal) == cache.maxKeyCount
}

func (cache *LRUCache) GetCurrentKeyCount() int {
	return len(cache.KeyToKeyVal)
}
