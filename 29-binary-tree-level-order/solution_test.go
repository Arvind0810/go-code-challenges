package levelorder

import (
	"slices"
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

func equalLevels(a, b [][]int) bool {
	return slices.EqualFunc(a, b, func(x, y []int) bool { return slices.Equal(x, y) })
}

func TestLevelOrder(t *testing.T) {
	cases := []struct {
		name string
		tree []*int
		want [][]int
	}{
		{"leetcode example", []*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"empty tree", nil, [][]int{}},
		{"single node", []*int{ptr(1)}, [][]int{{1}}},
		{"complete tree", []*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{"left skewed", []*int{ptr(1), ptr(2), nil, ptr(3), nil, ptr(4)}, [][]int{{1}, {2}, {3}, {4}}},
		{"right skewed", []*int{ptr(1), nil, ptr(2), nil, ptr(3)}, [][]int{{1}, {2}, {3}}},
		{"zigzag shape", []*int{ptr(1), ptr(2), ptr(3), nil, ptr(4), ptr(5), nil}, [][]int{{1}, {2, 3}, {4, 5}}},
		{"negatives and zero", []*int{ptr(0), ptr(-1), ptr(-2), ptr(-3)}, [][]int{{0}, {-1, -2}, {-3}}},
		{"gaps in last level", []*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, nil, ptr(5)}, [][]int{{1}, {2, 3}, {4, 5}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := LevelOrder(buildTree(tc.tree))
			if !equalLevels(got, tc.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tc.want)
			}
		})
	}
}
