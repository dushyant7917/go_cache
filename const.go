package go_cache

import (
	"container/list"
)

type (
	AnyType         = interface{}
	DLL             = *list.List // doubly linked list
	DLLElement      = *list.Element
	KeyToDLLElement = map[AnyType](DLLElement)
)

type CacheType int

const (
	FIFO CacheType = iota
	LRU
	LFU
)
