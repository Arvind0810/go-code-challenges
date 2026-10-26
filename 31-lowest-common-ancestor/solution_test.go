package lowestcommonancestor

import (
	"fmt"
	"math"
	"testing"
)

const null = math.MinInt

func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if found := findNode(root.Left, val); found != nil {
		return found
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	classic := []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}

	tests := []struct {
		name string
		tree []int
		p, q int
		want int
	}{
		{name: "split at root", tree: classic, p: 5, q: 1, want: 3},
		{name: "p is ancestor of q", tree: classic, p: 5, q: 4, want: 5},
		{name: "both deep in left subtree", tree: classic, p: 6, q: 4, want: 5},
		{name: "deep left and right", tree: classic, p: 7, q: 8, want: 3},
		{name: "both in right subtree", tree: classic, p: 0, q: 8, want: 1},
		{name: "siblings deep", tree: classic, p: 7, q: 4, want: 2},
		{name: "same node", tree: classic, p: 2, q: 2, want: 2},
		{name: "two nodes root is ancestor", tree: []int{1, 2}, p: 1, q: 2, want: 1},
		{name: "left-skewed chain", tree: []int{1, 2, null, 3, null, 4}, p: 4, q: 2, want: 2},
		{name: "negative values", tree: []int{-1, -2, -3, -4, -5}, p: -4, q: -3, want: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.tree)
			p, q := findNode(root, tc.p), findNode(root, tc.q)
			want := findNode(root, tc.want)
			got := LowestCommonAncestor(root, p, q)
			if got != want {
				gotVal := "nil"
				if got != nil {
					gotVal = fmt.Sprint(got.Val)
				}
				t.Errorf("LowestCommonAncestor(p=%d, q=%d) = %s, want node %d", tc.p, tc.q, gotVal, tc.want)
			}
		})
	}
}
