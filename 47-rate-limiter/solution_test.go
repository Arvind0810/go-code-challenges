package ratelimiter

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func newFakeClock() *fakeClock {
	return &fakeClock{now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

type step struct {
	advance time.Duration
	n       int
	want    bool
}

func runSteps(t *testing.T, b *TokenBucket, clock *fakeClock, steps []step) {
	t.Helper()
	for i, s := range steps {
		clock.Advance(s.advance)
		var got bool
		if s.n == 1 {
			got = b.Allow()
		} else {
			got = b.AllowN(s.n)
		}
		if got != s.want {
			t.Fatalf("step %d (advance %v, AllowN(%d)) = %v, want %v", i, s.advance, s.n, got, s.want)
		}
	}
}

func TestTokenBucketSequences(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		rate     float64
		steps    []step
	}{
		{
			name:     "starts full then empties",
			capacity: 3, rate: 1,
			steps: []step{{0, 1, true}, {0, 1, true}, {0, 1, true}, {0, 1, false}},
		},
		{
			name:     "refills one token per second",
			capacity: 2, rate: 1,
			steps: []step{
				{0, 1, true}, {0, 1, true}, {0, 1, false},
				{time.Second, 1, true}, {0, 1, false},
			},
		},
		{
			name:     "fractional tokens accumulate",
			capacity: 1, rate: 2,
			steps: []step{
				{0, 1, true},
				{250 * time.Millisecond, 1, false},
				{250 * time.Millisecond, 1, true},
				{0, 1, false},
			},
		},
		{
			name:     "refill is capped at capacity",
			capacity: 3, rate: 10,
			steps: []step{
				{0, 3, true},
				{time.Hour, 1, true}, {0, 1, true}, {0, 1, true}, {0, 1, false},
			},
		},
		{
			name:     "AllowN is all or nothing",
			capacity: 10, rate: 5,
			steps: []step{
				{0, 4, true},
				{0, 7, false},
				{0, 6, true},
				{0, 1, false},
				{time.Second, 5, true},
				{0, 1, false},
			},
		},
		{
			name:     "AllowN larger than capacity never succeeds",
			capacity: 5, rate: 100,
			steps: []step{{0, 6, false}, {time.Hour, 6, false}, {0, 5, true}},
		},
		{
			name:     "zero rate never refills",
			capacity: 2, rate: 0,
			steps: []step{{0, 2, true}, {time.Hour, 1, false}},
		},
		{
			name:     "zero capacity never allows",
			capacity: 0, rate: 10,
			steps: []step{{0, 1, false}, {time.Minute, 1, false}},
		},
		{
			name:     "non-positive n is always allowed",
			capacity: 1, rate: 0,
			steps: []step{{0, 1, true}, {0, 0, true}, {0, -3, true}, {0, 1, false}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clock := newFakeClock()
			b := NewTokenBucket(tc.capacity, tc.rate, clock)
			runSteps(t, b, clock, tc.steps)
		})
	}
}

func TestTokenBucketClockGoingBackwards(t *testing.T) {
	clock := newFakeClock()
	b := NewTokenBucket(2, 1, clock)
	runSteps(t, b, clock, []step{
		{0, 2, true},
		{-10 * time.Second, 1, false},
		{-time.Second, 1, false},
		{20 * time.Second, 1, true},
	})
}

func TestTokenBucketNilClockUsesRealClock(t *testing.T) {
	b := NewTokenBucket(1, 0, nil)
	if !b.Allow() {
		t.Fatal("first Allow() on a full bucket = false, want true")
	}
	if b.Allow() {
		t.Fatal("second Allow() with zero refill rate = true, want false")
	}
}

func runWithTimeout(t *testing.T, timeout time.Duration, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Fatalf("concurrent test did not finish within %v (deadlock?)", timeout)
	}
}

func TestTokenBucketConcurrentExactCount(t *testing.T) {
	const (
		capacity   = 1000
		goroutines = 100
		perG       = 25
	)
	clock := newFakeClock()
	b := NewTokenBucket(capacity, 0, clock)
	var allowed atomic.Int64
	runWithTimeout(t, 5*time.Second, func() {
		var wg sync.WaitGroup
		for range goroutines {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range perG {
					if b.Allow() {
						allowed.Add(1)
					}
				}
			}()
		}
		wg.Wait()
	})
	if got := allowed.Load(); got != capacity {
		t.Fatalf("%d of %d concurrent Allow() calls succeeded, want exactly %d", got, goroutines*perG, capacity)
	}
}

func TestTokenBucketConcurrentWithRefill(t *testing.T) {
	const (
		capacity = 10
		rate     = 100.0
		ticks    = 100
		tick     = 10 * time.Millisecond
	)
	clock := newFakeClock()
	b := NewTokenBucket(capacity, rate, clock)
	var allowed atomic.Int64
	runWithTimeout(t, 5*time.Second, func() {
		stop := make(chan struct{})
		var wg sync.WaitGroup
		for range 8 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-stop:
						return
					default:
					}
					if b.AllowN(1) {
						allowed.Add(1)
					}
				}
			}()
		}
		for range ticks {
			clock.Advance(tick)
			time.Sleep(time.Millisecond)
		}
		close(stop)
		wg.Wait()
	})
	maxAllowed := int64(capacity + rate*float64(ticks)*tick.Seconds() + 1)
	got := allowed.Load()
	if got > maxAllowed {
		t.Fatalf("allowed %d requests, want at most %d (capacity + refilled tokens)", got, maxAllowed)
	}
	if got < capacity {
		t.Fatalf("allowed %d requests, want at least the initial capacity %d", got, capacity)
	}
}
