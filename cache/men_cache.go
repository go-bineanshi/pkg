package cache

import (
	"sync"
	"time"
)

type MemCache struct {
	data map[string]MemCacheItem
	mu   sync.RWMutex
	ttl  time.Duration
}

type MemCacheItem struct {
	value     interface{}
	timestamp time.Time
}

func NewMemCache(ttl time.Duration) *MemCache {
	return &MemCache{
		data: make(map[string]MemCacheItem),
		ttl:  ttl,
	}
}

func (c *MemCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if val, found := c.data[key]; found {
		if time.Since(val.timestamp) > c.ttl*time.Second {
			delete(c.data, key)
			return nil, false
		}
		return val.value, true
	}

	return nil, false
}

func (c *MemCache) Set(key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = MemCacheItem{
		value:     value,
		timestamp: time.Now(),
	}
	return nil
}

func (c *MemCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
	return
}
