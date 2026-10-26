package pipeline

import (
	"context"
	"runtime"
	"slices"
	"sync"
	"testing"
	"time"
)

const recvTimeout = 3 * time.Second

func collect(t *testing.T, ch <-chan int) []int {
	t.Helper()
	timeout := time.After(recvTimeout)
	got := []int{}
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-timeout:
			t.Fatalf("channel was not closed within %v (received %d values so far)", recvTimeout, len(got))
			return nil
		}
	}
}

func recvN(t *testing.T, ch <-chan int, n int) []int {
	t.Helper()
	timeout := time.After(recvTimeout)
	got := make([]int, 0, n)
	for len(got) < n {
		select {
		case v, ok := <-ch:
			if !ok {
				t.Fatalf("channel closed after %d values, want at least %d", len(got), n)
			}
			got = append(got, v)
		case <-timeout:
			t.Fatalf("received only %d of %d values within %v", len(got), n, recvTimeout)
		}
	}
	return got
}

func waitForGoroutines(t *testing.T, baseline int) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > baseline {
		if time.Now().After(deadline) {
			t.Fatalf("goroutine leak: %d goroutines still running, want <= %d", runtime.NumGoroutine(), baseline)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func seq(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = i + 1
	}
	return out
}

func squares(nums []int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = n * n
	}
	return out
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{name: "empty", nums: nil},
		{name: "single", nums: []int{42}},
		{name: "ordered", nums: []int{3, 1, 4, 1, 5, 9, 2, 6}},
		{name: "many", nums: seq(1000)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := collect(t, Generate(context.Background(), tc.nums...))
			want := tc.nums
			if want == nil {
				want = []int{}
			}
			if !slices.Equal(got, want) {
				t.Errorf("Generate(%v) emitted %v, want %v in order", tc.nums, got, want)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	nums := []int{-3, 0, 2, 5, 11}
	got := collect(t, Square(context.Background(), Generate(context.Background(), nums...)))
	if want := squares(nums); !slices.Equal(got, want) {
		t.Errorf("Square emitted %v, want %v in order", got, want)
	}
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	t.Run("three sources", func(t *testing.T) {
		got := collect(t, Merge(ctx, Generate(ctx, 1, 2, 3), Generate(ctx, 10, 20), Generate(ctx, 100)))
		slices.Sort(got)
		if want := []int{1, 2, 3, 10, 20, 100}; !slices.Equal(got, want) {
			t.Errorf("Merge emitted %v (sorted), want %v", got, want)
		}
	})
	t.Run("no sources closes immediately", func(t *testing.T) {
		if got := collect(t, Merge(ctx)); len(got) != 0 {
			t.Errorf("Merge() emitted %v, want nothing", got)
		}
	})
	t.Run("empty sources", func(t *testing.T) {
		if got := collect(t, Merge(ctx, Generate(ctx), Generate(ctx))); len(got) != 0 {
			t.Errorf("Merge(empty, empty) emitted %v, want nothing", got)
		}
	})
}

func TestFanOut(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		wantCalls int
	}{
		{name: "four workers", n: 4, wantCalls: 4},
		{name: "one worker", n: 1, wantCalls: 1},
		{name: "zero treated as one", n: 0, wantCalls: 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			in := Generate(ctx, seq(200)...)

			var mu sync.Mutex
			var seen []<-chan int
			stage := func(ctx context.Context, c <-chan int) <-chan int {
				mu.Lock()
				seen = append(seen, c)
				mu.Unlock()
				return Square(ctx, c)
			}

			outs := FanOut(ctx, in, tc.n, stage)
			if len(outs) != tc.wantCalls {
				t.Fatalf("FanOut returned %d channels, want %d", len(outs), tc.wantCalls)
			}
			mu.Lock()
			calls := len(seen)
			sameInput := !slices.ContainsFunc(seen, func(c <-chan int) bool { return c != in })
			mu.Unlock()
			if calls != tc.wantCalls {
				t.Fatalf("stage function called %d times, want %d", calls, tc.wantCalls)
			}
			if !sameInput {
				t.Fatalf("every stage must read from the same input channel passed to FanOut")
			}
			got := collect(t, Merge(ctx, outs...))
			slices.Sort(got)
			if want := squares(seq(200)); !slices.Equal(got, want) {
				t.Errorf("fan-out/fan-in produced %d values, want the %d squares of 1..200", len(got), len(want))
			}
		})
	}
}

func TestFullPipeline(t *testing.T) {
	ctx := context.Background()
	nums := seq(500)
	out := Merge(ctx, FanOut(ctx, Generate(ctx, nums...), 8, Square)...)
	got := collect(t, out)
	slices.Sort(got)
	if want := squares(nums); !slices.Equal(got, want) {
		t.Errorf("pipeline produced %d values, want the %d squares of 1..%d", len(got), len(want), len(nums))
	}
}

func TestCancellationClosesOutput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	const total = 100000
	out := Merge(ctx, FanOut(ctx, Generate(ctx, seq(total)...), 4, Square)...)
	recvN(t, out, 10)
	cancel()
	rest := collect(t, out)
	if 10+len(rest) >= total {
		t.Errorf("received all %d values despite cancellation; stages must stop on ctx.Done()", total)
	}
}

func TestCancellationWithoutDrainingDoesNotLeak(t *testing.T) {
	baseline := runtime.NumGoroutine()
	ctx, cancel := context.WithCancel(context.Background())
	out := Merge(ctx, FanOut(ctx, Generate(ctx, seq(100000)...), 4, Square)...)
	recvN(t, out, 5)
	cancel()
	waitForGoroutines(t, baseline)
}

func TestCompletedPipelineDoesNotLeak(t *testing.T) {
	baseline := runtime.NumGoroutine()
	for range 20 {
		ctx := context.Background()
		collect(t, Merge(ctx, FanOut(ctx, Generate(ctx, seq(50)...), 3, Square)...))
	}
	waitForGoroutines(t, baseline)
}
