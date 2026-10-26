package serializetree

import (
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

func sameTree(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func TestSerializeDeserializeRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		tree []int
	}{
		{name: "empty tree", tree: nil},
		{name: "single node", tree: []int{1}},
		{name: "classic example", tree: []int{1, 2, 3, null, null, 4, 5}},
		{name: "complete tree", tree: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "left-skewed", tree: []int{1, 2, null, 3, null, 4}},
		{name: "right-skewed", tree: []int{1, null, 2, null, 3, null, 4}},
		{name: "negative and zero values", tree: []int{0, -1, -1000, null, -7}},
		{name: "multi-digit and duplicates", tree: []int{1000, 1000, -1000, 42, null, null, 42}},
		{name: "same values different shape", tree: []int{1, 1, null, null, 1}},
		{name: "zigzag", tree: []int{5, 4, null, null, 3, 2, null, null, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := buildTree(tc.tree)
			encoded := Serialize(original)
			decoded := Deserialize(encoded)
			if !sameTree(original, decoded) {
				t.Errorf("round trip failed for %s: Serialize produced %q, Deserialize rebuilt %d node(s), want a tree identical to the original %d node(s)",
					tc.name, encoded, countNodes(decoded), countNodes(original))
			}
		})
	}
}

func TestSerializeIsStable(t *testing.T) {
	tests := []struct {
		name string
		tree []int
	}{
		{name: "classic example", tree: []int{1, 2, 3, null, null, 4, 5}},
		{name: "right-skewed", tree: []int{1, null, 2, null, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			first := Serialize(buildTree(tc.tree))
			second := Serialize(Deserialize(first))
			if first != second {
				t.Errorf("Serialize(Deserialize(s)) = %q, want the same string %q", second, first)
			}
			if first == Serialize(nil) {
				t.Errorf("non-empty tree serialized to %q, the same as the empty tree", first)
			}
		})
	}
}
