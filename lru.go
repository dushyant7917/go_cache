package go_cache

type LRUCache struct {
	CapacityDetails
}

func (cache *LRUCache) Set(key, value interface{}) {
}

func (cache *LRUCache) Get(key interface{}) interface{} {
	return nil
}

func (cache *LRUCache) Evict() {
}

func (cache *LRUCache) Full() bool {
	return true
}

func (cache *LRUCache) GetCurrentKeyCount() int {
	return 0
}
