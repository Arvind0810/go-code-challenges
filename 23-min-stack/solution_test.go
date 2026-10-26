package minstack

import (
	"math"
	"testing"
)

type stackOp struct {
	name string
	arg  int
	want int
}

func opPush(v int) stackOp      { return stackOp{name: "push", arg: v} }
func opPop() stackOp            { return stackOp{name: "pop"} }
func opTop(want int) stackOp    { return stackOp{name: "top", want: want} }
func opGetMin(want int) stackOp { return stackOp{name: "getMin", want: want} }

func TestMinStack(t *testing.T) {
	cases := []struct {
		name string
		ops  []stackOp
	}{
		{"leetcode example", []stackOp{opPush(-2), opPush(0), opPush(-3), opGetMin(-3), opPop(), opTop(0), opGetMin(-2)}},
		{"single element", []stackOp{opPush(42), opTop(42), opGetMin(42)}},
		{"increasing pushes", []stackOp{opPush(1), opPush(2), opPush(3), opGetMin(1), opTop(3), opPop(), opGetMin(1), opTop(2)}},
		{"decreasing pushes", []stackOp{opPush(3), opPush(2), opPush(1), opGetMin(1), opPop(), opGetMin(2), opPop(), opGetMin(3)}},
		{"duplicate minimums", []stackOp{opPush(0), opPush(1), opPush(0), opGetMin(0), opPop(), opGetMin(0), opPop(), opGetMin(0)}},
		{"repeated min popped one by one", []stackOp{opPush(2), opPush(2), opPush(2), opPop(), opGetMin(2), opPop(), opGetMin(2), opTop(2)}},
		{"empty then reuse", []stackOp{opPush(5), opPop(), opPush(7), opPush(6), opGetMin(6), opTop(6), opPop(), opGetMin(7)}},
		{"extreme values", []stackOp{opPush(math.MaxInt64), opPush(math.MinInt64), opGetMin(math.MinInt64), opPop(), opGetMin(math.MaxInt64), opTop(math.MaxInt64)}},
		{"min returns after pop", []stackOp{opPush(5), opPush(3), opPush(4), opPush(1), opGetMin(1), opPop(), opGetMin(3), opPop(), opGetMin(3), opPop(), opGetMin(5)}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewMinStack()
			if s == nil {
				t.Fatal("NewMinStack() returned nil")
			}
			for i, o := range tc.ops {
				switch o.name {
				case "push":
					s.Push(o.arg)
				case "pop":
					s.Pop()
				case "top":
					if got := s.Top(); got != o.want {
						t.Fatalf("op #%d Top() = %d, want %d", i, got, o.want)
					}
				case "getMin":
					if got := s.GetMin(); got != o.want {
						t.Fatalf("op #%d GetMin() = %d, want %d", i, got, o.want)
					}
				}
			}
		})
	}
}
