package cache

import (
	"log"
	"sync"
	"time"
)

type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration int64
}

type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(duration).Unix()
	c.items[key] = CacheItem{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || item.Expiration <= time.Now().Unix() {
		return nil, false
	}
	if found {
		log.Printf("Cache item found: Key=%s, Value=%v, Expiration=%d", item.Key, item.Value, item.Expiration)
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
