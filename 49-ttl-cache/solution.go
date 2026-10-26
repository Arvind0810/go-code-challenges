package ttlcache

import "time"

// Clock abstracts time so expiry can be tested deterministically.
type Clock interface {
	Now() time.Time
}

// RealClock is a Clock backed by time.Now.
type RealClock struct{}

// Now returns the current wall-clock time.
func (RealClock) Now() time.Time { return time.Now() }

// Cache is a goroutine-safe key/value store whose entries expire after a
// time-to-live.
type Cache[K comparable, V any] struct {
	// TODO: implement
}

// New returns an empty cache. Entries stored with Set expire defaultTTL
// after they are written; a defaultTTL <= 0 means they never expire. A nil
// clock means RealClock.
func New[K comparable, V any](defaultTTL time.Duration, clock Clock) *Cache[K, V] {
	// TODO: implement
	return &Cache[K, V]{}
}

// Set stores v under k using the cache's default TTL, replacing any existing
// value and resetting its expiry.
func (c *Cache[K, V]) Set(k K, v V) {
	// TODO: implement
}

// SetWithTTL stores v under k with its own ttl. A ttl <= 0 means the entry
// never expires.
func (c *Cache[K, V]) SetWithTTL(k K, v V, ttl time.Duration) {
	// TODO: implement
}

// Get returns the value for k and true if it exists and has not expired. An
// entry is expired once the clock reaches its expiry time; expired entries
// found by Get are removed.
func (c *Cache[K, V]) Get(k K) (V, bool) {
	// TODO: implement
	var zero V
	return zero, false
}

// Delete removes k from the cache. Deleting a missing key is a no-op.
func (c *Cache[K, V]) Delete(k K) {
	// TODO: implement
}

// Len returns the number of entries currently stored, including expired
// entries that have not yet been evicted by Get or the janitor.
func (c *Cache[K, V]) Len() int {
	// TODO: implement
	return 0
}

// StartJanitor starts a background goroutine that evicts all expired entries
// every interval until Stop is called. Calling it while a janitor is already
// running has no effect.
func (c *Cache[K, V]) StartJanitor(interval time.Duration) {
	// TODO: implement
}

// Stop stops the janitor and waits for its goroutine to exit. It is safe to
// call Stop more than once, or without having started a janitor.
func (c *Cache[K, V]) Stop() {
	// TODO: implement
}
