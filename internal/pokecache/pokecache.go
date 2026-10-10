package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
			entries: map[string]cacheEntry{},
			mu: &sync.RWMutex{},
	}
	go cache.reaploob(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	entry, exists := c.entries[key] 
	if !exists {
		return []byte{}, exists
	}
	return entry.val, exists
}

func (c *Cache) reaploob(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		for key, val := range c.entries {
			c.mu.Lock()
			
			if val.createdAt.Before(time.Now().UTC().Add(-interval)) {
				delete(c.entries, key)
			}
			c.mu.Unlock()
		}
	}

}
