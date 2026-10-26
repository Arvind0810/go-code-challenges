package pipeline

import "context"

// Generate returns a channel that emits nums in order and is closed after
// the last value or as soon as ctx is done.
func Generate(ctx context.Context, nums ...int) <-chan int {
	// TODO: implement
	return closedChan()
}

// Square returns a channel that emits the square of every value received
// from in, in the same order. It is closed when in is closed or ctx is done.
func Square(ctx context.Context, in <-chan int) <-chan int {
	// TODO: implement
	return closedChan()
}

// FanOut starts n copies of the stage fn, all reading from the shared input
// channel in, and returns their n output channels. A value of n below 1 is
// treated as 1.
func FanOut(ctx context.Context, in <-chan int, n int, fn func(context.Context, <-chan int) <-chan int) []<-chan int {
	// TODO: implement
	return nil
}

// Merge returns a single channel that emits every value received from chans
// and is closed once all of them are closed or ctx is done.
func Merge(ctx context.Context, chans ...<-chan int) <-chan int {
	// TODO: implement
	return closedChan()
}

func closedChan() <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}
