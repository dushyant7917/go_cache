package go_cache

type FIFOCache struct {
	CapacityDetails
	KeyValPairs DLL
	KeyToKeyVal KeyToDLLElement
}

func (cache *FIFOCache) Set(key, val AnyType) {
	if kv, keyFound := cache.KeyToKeyVal[key]; keyFound {
		keyVal := kv.Value.(*KeyVal)
		keyVal.val = val
		return
	}

	if cacheFull(cache.maxKeyCount, cache.KeyToKeyVal) {
		cache.Evict()
	}
	cache.KeyValPairs.PushBack(&KeyVal{key: key, val: val})
	cache.KeyToKeyVal[key] = cache.KeyValPairs.Back()
}

func (cache *FIFOCache) Get(key AnyType) AnyType {
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
