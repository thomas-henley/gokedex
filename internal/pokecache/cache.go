package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mtx      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries:  map[string]cacheEntry{},
		interval: interval,
		mtx:      &sync.Mutex{},
	}
	cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mtx.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	entry, found := c.entries[key]
	c.mtx.Unlock()
	if found {
		return entry.val, true
	}
	return nil, false
}

func (c Cache) reapLoop() {
	reaper := time.NewTicker(c.interval)

	go func() {
		for {
			<-reaper.C
			c.mtx.Lock()
			fmt.Println("Reaping cache entries...")
			for k, v := range c.entries {
				if time.Since(v.createdAt) > c.interval {
					delete(c.entries, k)
				}
			}
			c.mtx.Unlock()
		}
	}()
}
