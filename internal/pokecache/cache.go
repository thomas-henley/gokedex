package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mtx     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mtx:     &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	reaper := time.NewTicker(interval)

	for range reaper.C {
		c.reap(interval)

	}
}

func (c Cache) reap(interval time.Duration) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for k, v := range c.entries {
		if time.Since(v.createdAt) > interval {
			delete(c.entries, k)
		}
	}
}
