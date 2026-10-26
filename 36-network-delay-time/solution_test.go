package networkdelay

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n, k  int
		want  int
	}{
		{name: "classic example", times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, n: 4, k: 2, want: 2},
		{name: "single edge reachable", times: [][]int{{1, 2, 1}}, n: 2, k: 1, want: 1},
		{name: "edges are directed", times: [][]int{{1, 2, 1}}, n: 2, k: 2, want: -1},
		{name: "single node", times: nil, n: 1, k: 1, want: 0},
		{name: "indirect path is shorter", times: [][]int{{1, 2, 10}, {1, 3, 1}, {3, 2, 2}}, n: 3, k: 1, want: 3},
		{name: "disconnected node", times: [][]int{{1, 2, 1}, {2, 3, 1}}, n: 4, k: 1, want: -1},
		{name: "zero weight edges", times: [][]int{{1, 2, 0}, {2, 3, 0}}, n: 3, k: 1, want: 0},
		{name: "parallel edges take cheapest", times: [][]int{{1, 2, 5}, {1, 2, 2}, {2, 3, 4}}, n: 3, k: 1, want: 6},
		{name: "cycle in graph", times: [][]int{{1, 2, 1}, {2, 3, 2}, {3, 1, 4}, {3, 4, 3}}, n: 4, k: 2, want: 6},
		{
			name:  "greedy first edge is wrong",
			times: [][]int{{1, 2, 1}, {2, 5, 100}, {1, 3, 2}, {3, 4, 2}, {4, 5, 2}},
			n:     5, k: 1, want: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NetworkDelayTime(tc.times, tc.n, tc.k)
			if got != tc.want {
				t.Errorf("NetworkDelayTime(%v, n=%d, k=%d) = %d, want %d", tc.times, tc.n, tc.k, got, tc.want)
			}
		})
	}
}
