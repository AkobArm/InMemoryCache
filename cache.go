package InMemoryCache

import (
	"errors"
	"sync"
	"time"
)

type Item struct {
	Value     interface{}
	ExpiresAt time.Time
}

type Cache struct {
	data map[string]*Item
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*Item),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = &Item{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	if !ok || item.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("key not found")
	}
	return item.Value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]*Item)
}

func (c *Cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) Values() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	values := make([]interface{}, 0, len(c.data))
	for _, item := range c.data {
		if item.ExpiresAt.After(time.Now()) {
			values = append(values, item.Value)
		}
	}
	return values
}

func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	return ok && item.ExpiresAt.After(time.Now())
}
