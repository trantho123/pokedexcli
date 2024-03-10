package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val      []byte
	createAt time.Time
}

func NewCache(internal time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(internal)
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:      val,
		createAt: time.Now().UTC(),
	}
}

func (c *Cache) reapLoop(internal time.Duration) {
	ticker := time.NewTicker(internal)
	for range ticker.C {
		c.reap(internal)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
