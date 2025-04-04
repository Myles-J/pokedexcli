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
	mu       sync.RWMutex
	interval time.Duration
	cache    map[string]CacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		cache:    make(map[string]CacheEntry),
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cache[key] = cacheEntry

	return cacheEntry.val, false
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

// reapLoop is a helper function that deletes expired entries from the cache
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
