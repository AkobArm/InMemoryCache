# InMemoryCache

The project is an implementation of an in-memory cache in Go.

## To use the package, run the command:

```bash
go get -u github.com/AkobArm/InMemoryCache
```

## Basic methods

- `Set(key string, value interface{}, ttl time.Duration)` - writing the value `value` to the cache using the key `key` with lifetime `ttl` 
- `Get(key string)` - getting the value from the cache by key `key` 
- `Delete(key string)` - deleting a value from the cache by key `key` 
- `Clear()` - deleting all values from the cache 
- `Keys()` - getting all keys from the cache 
- `Values ()` - getting all values from the cache 
- `Exists(key string)` - checking if the key `key` is in the cache

## Usage example

```go
package main

import (
	"fmt"
	"InMemoryCache"
    "time"
)

func main() {
	cache := InMemoryCache.NewCache()
	cache.Set("key", "value", time.Minute)
	val, err := cache.Get("key")
	if err == nil {
		fmt.Println(val)
	}
	cache.Delete("key")
	cache.Clear()
}