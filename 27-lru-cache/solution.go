package lrucache

// LRUCache is a fixed-capacity key/value cache that evicts the least recently
// used entry when full. Add whatever fields you need.
type LRUCache struct{}

// NewLRUCache returns an empty cache that holds at most capacity entries.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{}
}

// Get returns the value for key and marks it as most recently used, or -1 if
// the key is not present.
func (c *LRUCache) Get(key int) int {
	// TODO: implement
	return 0
}

// Put inserts or updates key and marks it as most recently used, evicting the
// least recently used entry if the cache is over capacity.
func (c *LRUCache) Put(key, value int) {
	// TODO: implement
}
