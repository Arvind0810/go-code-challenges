package validatebst

import (
	"math"
	"testing"
)

func ptr(v int) *int { return &v }

// buildTree builds a tree from LeetCode-style level-order input, where nil
// marks a missing child.
func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestIsValidBST(t *testing.T) {
	cases := []struct {
		name string
		tree []*int
		want bool
	}{
		{"empty tree", nil, true},
		{"single node", []*int{ptr(1)}, true},
		{"simple valid", []*int{ptr(2), ptr(1), ptr(3)}, true},
		{"right child smaller than root", []*int{ptr(5), ptr(1), ptr(4), nil, nil, ptr(3), ptr(6)}, false},
		{"deep violation of ancestor bound", []*int{ptr(5), ptr(4), ptr(6), nil, nil, ptr(3), ptr(7)}, false},
		{"duplicate on the left", []*int{ptr(2), ptr(2)}, false},
		{"duplicate on the right", []*int{ptr(2), nil, ptr(2)}, false},
		{"all duplicates", []*int{ptr(2), ptr(2), ptr(2)}, false},
		{"larger valid tree", []*int{ptr(8), ptr(4), ptr(12), ptr(2), ptr(6), ptr(10), ptr(14), ptr(1), ptr(3), ptr(5), ptr(7)}, true},
		{"single MinInt64 node", []*int{ptr(math.MinInt64)}, true},
		{"single MaxInt64 node", []*int{ptr(math.MaxInt64)}, true},
		{"MinInt64 and MaxInt64 children", []*int{ptr(0), ptr(math.MinInt64), ptr(math.MaxInt64)}, true},
		{"MaxInt64 root with MinInt64 left", []*int{ptr(math.MaxInt64), ptr(math.MinInt64)}, true},
		{"duplicate MaxInt64 on the right", []*int{ptr(math.MaxInt64), nil, ptr(math.MaxInt64)}, false},
		{"duplicate MinInt64 on the left", []*int{ptr(math.MinInt64), ptr(math.MinInt64)}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValidBST(buildTree(tc.tree)); got != tc.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tc.want)
			}
		})
	}
}
