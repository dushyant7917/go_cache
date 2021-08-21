package go_cache

type LFUCache struct {
	CapacityDetails
}

func (cache *LFUCache) Set(key, value interface{}) {
}

func (cache *LFUCache) Get(key interface{}) interface{} {
	return nil
}

func (cache *LFUCache) Evict() {
}

func (cache *LFUCache) Full() bool {
	return false
}

func (cache *LFUCache) GetCurrentKeyCount() int {
	return 0
}
