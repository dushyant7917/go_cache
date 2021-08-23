package go_cache

func moveKVToBack(cache *LRUCache, keyVal DLLElement, key, val AnyType) {
	cache.KeyValPairs.Remove(keyVal)
	cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
	cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
}

type LRUCache struct {
	CapacityDetails
	KeyValPairs DLL
	KeyToKeyVal KeyToDLLElement
}

func (cache *LRUCache) Set(key, val AnyType) {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		moveKVToBack(cache, kv, key, val)
		return
	}

	if cacheFull(cache.maxKeyCount, cache.KeyToKeyVal) {
		cache.Evict()
	}
	cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
	cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
}

func (cache *LRUCache) Get(key AnyType) AnyType {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		val := kv.Value.(*KeyVal).val
		moveKVToBack(cache, kv, key, val)
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
