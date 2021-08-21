package go_cache

type CacheType int

const (
	FIFO CacheType = iota
	LRU
	LFU
)
