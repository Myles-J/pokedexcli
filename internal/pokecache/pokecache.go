package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.RWMutex
	cache map[string]CacheEntry
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		interval: interval,
		cache:    make(map[string]CacheEntry),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	if time.Since(entry.createdAt) > c.interval {
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.cache, key)
		return nil, false
	}

	return entry.val, true
}
