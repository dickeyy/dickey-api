package cache

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Item represents a cached item with expiration
type Item struct {
	Value      interface{}
	Expiration int64
	CreatedAt  time.Time
}

// Cache represents an in-memory cache with expiration
type Cache struct {
	items        map[string]Item
	mu           sync.RWMutex
	cleanupTimer *time.Ticker
	name         string
}

// New creates a new cache instance
func New() *Cache {
	return NewWithName("default")
}

// NewWithName creates a new cache instance with a name for logging
func NewWithName(name string) *Cache {
	cache := &Cache{
		items:        make(map[string]Item),
		cleanupTimer: time.NewTicker(time.Second),
		name:         name,
	}

	// Start a goroutine to clean up expired items
	go cache.startCleanupTimer()

	log.WithFields(log.Fields{
		"cache": name,
	}).Info("Cache initialized")

	return cache
}

// Set adds an item to the cache with a 10-second expiration
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	// Set expiration to 10 seconds from now
	expiration := now.Add(10 * time.Second).UnixNano()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		CreatedAt:  now,
	}

	log.WithFields(log.Fields{
		"cache": c.name,
		"key":   key,
	}).Debug("Item added to cache")
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		log.WithFields(log.Fields{
			"cache": c.name,
			"key":   key,
		}).Debug("Cache miss")
		return nil, false
	}

	// Check if the item has expired
	now := time.Now().UnixNano()
	if now > item.Expiration {
		log.WithFields(log.Fields{
			"cache": c.name,
			"key":   key,
			"age":   time.Since(item.CreatedAt).String(),
		}).Debug("Cache item expired")
		return nil, false
	}

	log.WithFields(log.Fields{
		"cache": c.name,
		"key":   key,
		"age":   time.Since(item.CreatedAt).String(),
	}).Debug("Cache hit")
	return item.Value, true
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[key]; exists {
		delete(c.items, key)
		log.WithFields(log.Fields{
			"cache": c.name,
			"key":   key,
		}).Debug("Item deleted from cache")
	}
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]Item)
	log.WithFields(log.Fields{
		"cache": c.name,
	}).Info("Cache cleared")
}

// Count returns the number of items in the cache
func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// startCleanupTimer starts a timer to clean up expired items
func (c *Cache) startCleanupTimer() {
	for {
		<-c.cleanupTimer.C
		c.cleanup()
	}
}

// cleanup removes expired items from the cache
func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()
	expiredCount := 0

	for k, v := range c.items {
		if now > v.Expiration {
			delete(c.items, k)
			expiredCount++
		}
	}

	if expiredCount > 0 {
		log.WithFields(log.Fields{
			"cache":        c.name,
			"expiredCount": expiredCount,
			"remaining":    len(c.items),
		}).Debug("Expired items removed from cache")
	}
}

// Close stops the cleanup timer
func (c *Cache) Close() {
	c.cleanupTimer.Stop()
	log.WithFields(log.Fields{
		"cache": c.name,
	}).Info("Cache closed")
}
