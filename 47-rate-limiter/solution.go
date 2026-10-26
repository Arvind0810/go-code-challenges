package ratelimiter

import "time"

// Clock abstracts time so the limiter can be tested deterministically.
type Clock interface {
	Now() time.Time
}

// RealClock is a Clock backed by time.Now.
type RealClock struct{}

// Now returns the current wall-clock time.
func (RealClock) Now() time.Time { return time.Now() }

// TokenBucket is a goroutine-safe token-bucket rate limiter.
type TokenBucket struct {
	// TODO: implement
}

// NewTokenBucket returns a bucket that holds at most capacity tokens, starts
// full, and refills continuously at refillPerSecond tokens per second. A nil
// clock means RealClock.
func NewTokenBucket(capacity int, refillPerSecond float64, clock Clock) *TokenBucket {
	// TODO: implement
	return &TokenBucket{}
}

// Allow reports whether one token is available and, if so, consumes it.
func (b *TokenBucket) Allow() bool {
	// TODO: implement
	return false
}

// AllowN reports whether n tokens are available and, if so, consumes all of
// them. If fewer than n tokens are available, nothing is consumed. AllowN
// returns true without consuming anything when n <= 0.
func (b *TokenBucket) AllowN(n int) bool {
	// TODO: implement
	return false
}
