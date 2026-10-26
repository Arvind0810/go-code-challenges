package workerpool

import "context"

// Job is a unit of work submitted to the pool.
type Job struct {
	ID    int
	Input int
}

// Result is the outcome of processing one Job. RunPool always sets JobID to
// the ID of the job it belongs to and Err to the error returned by the work
// function (or the context error for jobs that were never started).
type Result struct {
	JobID  int
	Output int
	Err    error
}

// RunPool processes jobs using at most workers concurrent calls to fn and
// returns one Result per job, in the same order as jobs. A workers value
// below 1 is treated as 1. Once ctx is done, jobs that have not started are
// skipped and their Result.Err is ctx.Err(). RunPool returns only after
// every call to fn it started has returned.
func RunPool(ctx context.Context, workers int, jobs []Job, fn func(context.Context, Job) (Result, error)) []Result {
	// TODO: implement
	return nil
}
