package cache

import (
	"sync"
)

// LastFmCache is a singleton cache instance for Last.fm API responses
var (
	lastFmCache     *Cache
	lastFmCacheOnce sync.Once
)

// GetLastFmCache returns the singleton instance of the Last.fm cache
func GetLastFmCache() *Cache {
	lastFmCacheOnce.Do(func() {
		lastFmCache = NewWithName("lastfm")
	})
	return lastFmCache
}

// GenerateLastFmCacheKey generates a cache key for Last.fm API requests
func GenerateLastFmCacheKey(method string, user string) string {
	return "lastfm:" + method + ":" + user
}
