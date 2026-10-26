package lrucache

import "testing"

type cacheOp struct {
	get   bool
	key   int
	value int
	want  int
}

func opPut(key, value int) cacheOp { return cacheOp{key: key, value: value} }
func opGet(key, want int) cacheOp  { return cacheOp{get: true, key: key, want: want} }

func TestLRUCache(t *testing.T) {
	cases := []struct {
		name     string
		capacity int
		ops      []cacheOp
	}{
		{"leetcode example", 2, []cacheOp{
			opPut(1, 1), opPut(2, 2), opGet(1, 1), opPut(3, 3), opGet(2, -1),
			opPut(4, 4), opGet(1, -1), opGet(3, 3), opGet(4, 4),
		}},
		{"get on empty cache", 3, []cacheOp{opGet(1, -1), opGet(0, -1)}},
		{"capacity one evicts immediately", 1, []cacheOp{
			opPut(1, 10), opGet(1, 10), opPut(2, 20), opGet(1, -1), opGet(2, 20),
		}},
		{"update existing key does not evict", 2, []cacheOp{
			opPut(1, 1), opPut(2, 2), opPut(1, 100), opGet(1, 100), opGet(2, 2),
		}},
		{"update refreshes recency", 2, []cacheOp{
			opPut(1, 1), opPut(2, 2), opPut(1, 11), opPut(3, 3),
			opGet(2, -1), opGet(1, 11), opGet(3, 3),
		}},
		{"get refreshes recency", 3, []cacheOp{
			opPut(1, 1), opPut(2, 2), opPut(3, 3), opGet(1, 1), opPut(4, 4),
			opGet(2, -1), opGet(1, 1), opGet(3, 3), opGet(4, 4),
		}},
		{"missing get does not affect order", 2, []cacheOp{
			opPut(1, 1), opPut(2, 2), opGet(9, -1), opPut(3, 3), opGet(1, -1), opGet(2, 2),
		}},
		{"zero key and zero value", 2, []cacheOp{
			opPut(0, 0), opGet(0, 0), opPut(1, 1), opPut(2, 2), opGet(0, -1),
		}},
		{"many evictions in a row", 2, []cacheOp{
			opPut(1, 1), opPut(2, 2), opPut(3, 3), opPut(4, 4), opPut(5, 5),
			opGet(1, -1), opGet(2, -1), opGet(3, -1), opGet(4, 4), opGet(5, 5),
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewLRUCache(tc.capacity)
			if c == nil {
				t.Fatal("NewLRUCache() returned nil")
			}
			for i, o := range tc.ops {
				if !o.get {
					c.Put(o.key, o.value)
					continue
				}
				if got := c.Get(o.key); got != o.want {
					t.Fatalf("op #%d Get(%d) = %d, want %d", i, o.key, got, o.want)
				}
			}
		})
	}
}

func TestLRUCacheLarge(t *testing.T) {
	const capacity = 1000
	c := NewLRUCache(capacity)
	for i := range 3 * capacity {
		c.Put(i, i*2)
	}
	for i := range 2 * capacity {
		if got := c.Get(i); got != -1 {
			t.Fatalf("Get(%d) = %d, want -1 (should have been evicted)", i, got)
		}
	}
	for i := 2 * capacity; i < 3*capacity; i++ {
		if got := c.Get(i); got != i*2 {
			t.Fatalf("Get(%d) = %d, want %d", i, got, i*2)
		}
	}
}
