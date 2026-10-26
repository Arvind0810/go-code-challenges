package workerpool

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

const callTimeout = 5 * time.Second

type workFn = func(context.Context, Job) (Result, error)

func runPool(t *testing.T, ctx context.Context, workers int, jobs []Job, fn workFn) []Result {
	t.Helper()
	done := make(chan []Result, 1)
	go func() { done <- RunPool(ctx, workers, jobs, fn) }()
	select {
	case res := <-done:
		if len(res) != len(jobs) {
			t.Fatalf("RunPool returned %d results, want %d (one per job)", len(res), len(jobs))
		}
		for i, r := range res {
			if r.JobID != jobs[i].ID {
				t.Fatalf("results[%d].JobID = %d, want %d (results must be in input order)", i, r.JobID, jobs[i].ID)
			}
		}
		return res
	case <-time.After(callTimeout):
		t.Fatalf("RunPool did not return within %v (deadlock or missing wg.Wait/close?)", callTimeout)
		return nil
	}
}

func makeJobs(n int) []Job {
	jobs := make([]Job, n)
	for i := range jobs {
		jobs[i] = Job{ID: 1000 + i, Input: i}
	}
	return jobs
}

type concurrencyTracker struct {
	current atomic.Int64
	peak    atomic.Int64
	calls   atomic.Int64
}

func (c *concurrencyTracker) enter() {
	c.calls.Add(1)
	n := c.current.Add(1)
	for {
		p := c.peak.Load()
		if n <= p || c.peak.CompareAndSwap(p, n) {
			return
		}
	}
}

func (c *concurrencyTracker) exit() { c.current.Add(-1) }

func waitForGoroutines(t *testing.T, baseline int) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > baseline {
		if time.Now().After(deadline) {
			t.Fatalf("goroutine leak: %d goroutines running, want <= %d", runtime.NumGoroutine(), baseline)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestRunPoolPreservesOrder(t *testing.T) {
	jobs := makeJobs(40)
	fn := func(_ context.Context, j Job) (Result, error) {
		time.Sleep(time.Duration(40-j.Input) * 200 * time.Microsecond)
		return Result{Output: j.Input * j.Input}, nil
	}
	res := runPool(t, context.Background(), 8, jobs, fn)
	for i, r := range res {
		if r.Err != nil {
			t.Errorf("results[%d].Err = %v, want nil", i, r.Err)
		}
		if want := jobs[i].Input * jobs[i].Input; r.Output != want {
			t.Errorf("results[%d].Output = %d, want %d", i, r.Output, want)
		}
	}
}

func TestRunPoolCapturesErrorsPerJob(t *testing.T) {
	errOdd := errors.New("odd input")
	jobs := makeJobs(20)
	fn := func(_ context.Context, j Job) (Result, error) {
		if j.Input%2 == 1 {
			return Result{}, fmt.Errorf("job %d: %w", j.ID, errOdd)
		}
		return Result{Output: j.Input + 1}, nil
	}
	res := runPool(t, context.Background(), 4, jobs, fn)
	for i, r := range res {
		if jobs[i].Input%2 == 1 {
			if !errors.Is(r.Err, errOdd) {
				t.Errorf("results[%d].Err = %v, want an error wrapping %v", i, r.Err, errOdd)
			}
			continue
		}
		if r.Err != nil || r.Output != jobs[i].Input+1 {
			t.Errorf("results[%d] = {Output:%d Err:%v}, want {Output:%d Err:<nil>}", i, r.Output, r.Err, jobs[i].Input+1)
		}
	}
}

func TestRunPoolMaxConcurrency(t *testing.T) {
	tests := []struct {
		name     string
		workers  int
		jobs     int
		wantPeak int64
	}{
		{name: "three workers", workers: 3, jobs: 30, wantPeak: 3},
		{name: "more workers than jobs", workers: 10, jobs: 4, wantPeak: 4},
		{name: "single worker", workers: 1, jobs: 10, wantPeak: 1},
		{name: "zero workers treated as one", workers: 0, jobs: 5, wantPeak: 1},
		{name: "negative workers treated as one", workers: -2, jobs: 5, wantPeak: 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var tr concurrencyTracker
			fn := func(_ context.Context, j Job) (Result, error) {
				tr.enter()
				defer tr.exit()
				time.Sleep(20 * time.Millisecond)
				return Result{Output: j.Input}, nil
			}
			runPool(t, context.Background(), tc.workers, makeJobs(tc.jobs), fn)
			if got := tr.calls.Load(); got != int64(tc.jobs) {
				t.Errorf("fn called %d times, want %d", got, tc.jobs)
			}
			if got := tr.peak.Load(); got != tc.wantPeak {
				t.Errorf("peak concurrency = %d, want %d", got, tc.wantPeak)
			}
		})
	}
}

func TestRunPoolEmptyJobs(t *testing.T) {
	var calls atomic.Int64
	fn := func(context.Context, Job) (Result, error) {
		calls.Add(1)
		return Result{}, nil
	}
	runPool(t, context.Background(), 4, nil, fn)
	if calls.Load() != 0 {
		t.Errorf("fn called %d times for no jobs, want 0", calls.Load())
	}
}

func TestRunPoolCancellation(t *testing.T) {
	baseline := runtime.NumGoroutine()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const total = 200
	var calls atomic.Int64
	fn := func(ctx context.Context, j Job) (Result, error) {
		if calls.Add(1) == 5 {
			cancel()
		}
		select {
		case <-time.After(10 * time.Millisecond):
			return Result{Output: j.Input}, nil
		case <-ctx.Done():
			return Result{}, ctx.Err()
		}
	}

	start := time.Now()
	res := runPool(t, ctx, 2, makeJobs(total), fn)
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("RunPool took %v after cancellation, want it to stop promptly", elapsed)
	}
	if got := calls.Load(); got > 10 {
		t.Errorf("fn called %d times, want at most ~%d (jobs must not start after ctx is cancelled)", got, 5+2)
	}

	canceled := 0
	for i, r := range res {
		switch {
		case r.Err == nil:
		case errors.Is(r.Err, context.Canceled):
			canceled++
		default:
			t.Errorf("results[%d].Err = %v, want nil or context.Canceled", i, r.Err)
		}
	}
	if canceled < total-10 {
		t.Errorf("%d results carry context.Canceled, want at least %d", canceled, total-10)
	}
	waitForGoroutines(t, baseline)
}

func TestRunPoolAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var calls atomic.Int64
	fn := func(context.Context, Job) (Result, error) {
		calls.Add(1)
		return Result{}, nil
	}
	res := runPool(t, ctx, 4, makeJobs(10), fn)
	if calls.Load() != 0 {
		t.Errorf("fn called %d times with an already-cancelled context, want 0", calls.Load())
	}
	for i, r := range res {
		if !errors.Is(r.Err, context.Canceled) {
			t.Errorf("results[%d].Err = %v, want context.Canceled", i, r.Err)
		}
	}
}

func TestRunPoolNoGoroutineLeak(t *testing.T) {
	baseline := runtime.NumGoroutine()
	fn := func(_ context.Context, j Job) (Result, error) { return Result{Output: j.Input}, nil }
	for range 20 {
		runPool(t, context.Background(), 8, makeJobs(50), fn)
	}
	waitForGoroutines(t, baseline)
}
