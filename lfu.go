package go_cache

import (
	"container/list"
)

type KeyValFreq struct {
	KeyVal
	freqElement DLLElement
}

type Freq struct {
	freq    int
	keyVals DLL
}

type LFUCache struct {
	CapacityDetails
	Frequencies     DLL
	KeyToKeyValFreq KeyToDLLElement
}

func (cache *LFUCache) Set(key, val AnyType) {
	if _, keyFound := cache.KeyToKeyValFreq[key]; keyFound {
		return
	}

	if cacheFull(cache.maxKeyCount, cache.KeyToKeyValFreq) {
		cache.Evict()
	}
	frontFreq := cache.Frequencies.Front()
	if frontFreq == nil || frontFreq.Value.(*Freq).freq != 1 {
		cache.Frequencies.PushFront(&Freq{freq: 1, keyVals: list.New()})
	}
	frontFreq = cache.Frequencies.Front()
	kvf := &KeyValFreq{KeyVal: KeyVal{key: key, val: val}, freqElement: frontFreq}
	frontFreq.Value.(*Freq).keyVals.PushBack(kvf)
	cache.KeyToKeyValFreq[key] = frontFreq.Value.(*Freq).keyVals.Back()
}

func (cache *LFUCache) Get(key AnyType) AnyType {
	if _, keyFound := cache.KeyToKeyValFreq[key]; !keyFound {
		return nil
	}

	kvfElement := cache.KeyToKeyValFreq[key]
	val := kvfElement.Value.(*KeyValFreq).val
	freqElement := kvfElement.Value.(*KeyValFreq).freqElement
	freqElement.Value.(*Freq).keyVals.Remove(kvfElement)
	freq := freqElement.Value.(*Freq).freq
	nextFreqElement := freqElement.Next()
	if nextFreqElement == nil || nextFreqElement.Value.(*Freq).freq != freq+1 {
		newFreqElement := &Freq{freq: freq + 1, keyVals: list.New()}
		cache.Frequencies.InsertAfter(newFreqElement, freqElement)
	}
	prvFreqElement := freqElement
	freqElement = freqElement.Next()
	if prvFreqElement.Value.(*Freq).keyVals.Len() == 0 {
		cache.Frequencies.Remove(prvFreqElement)
	}
	kvf := &KeyValFreq{KeyVal: KeyVal{key: key, val: val}, freqElement: freqElement}
	freqElement.Value.(*Freq).keyVals.PushBack(kvf)
	cache.KeyToKeyValFreq[key] = freqElement.Value.(*Freq).keyVals.Back()
	return val
}

func (cache *LFUCache) Evict() {
	if len(cache.KeyToKeyValFreq) == 0 {
		return
	}
	leastFreqElement := cache.Frequencies.Front()
	first_kvfElement := leastFreqElement.Value.(*Freq).keyVals.Front()
	key := first_kvfElement.Value.(*KeyValFreq).key
	leastFreqElement.Value.(*Freq).keyVals.Remove(first_kvfElement)
	delete(cache.KeyToKeyValFreq, key)
}
