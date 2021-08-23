package go_cache

func cacheFull(maxKeyCount int, keyToItem KeyToDLLElement) bool {
	return len(keyToItem) == maxKeyCount
}
