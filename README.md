# In-memory cache

#### Types of eviction policies supported:
1. FIFO
2. LRU
3. LFU

#### Runing all the unit tests: 
`go test`

#### `go_cache` can be used as library in other go programs (see below example).
```
package main

import (
	"fmt"

	"github.com/dushyant7917/go_cache"
)

func main() {
	cache := go_cache.NewCache(2, go_cache.LRU)
	// cache := go_cache.NewCache(5, go_cache.FIFO)
	// cache := go_cache.NewCache(1, go_cache.LFU)
	cache.Set(1, "abc")
	fmt.Println(cache.Get(1))
	cache.Set("foo", 45.67)
	fmt.Println(cache.Get("foo"))
}

```
#### Todo:
1. Adding mulithread support (currently no locking is used on the data structures).
2. Handling more edge cases and exceptions.
3. Adding benchmark tests.