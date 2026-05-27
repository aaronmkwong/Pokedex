package pokecache

import (
	"sync"
	"time"
)

// cacheEntry stores a cached value and when it was created
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache stores cached entries and protects them with a mutex
type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

// NewCache creates a new cache and starts the background reap loop
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}

	// Start background cleanup goroutine
	go c.reapLoop(interval)

	return c
}

// Add stores a value in the cache with the current timestamp
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a value from the cache
// Returns the value and true if found, otherwise nil and false
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

// reapLoop periodically removes expired cache entries
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

// reap removes entries older than the provided interval
func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cache {
		if time.Since(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}