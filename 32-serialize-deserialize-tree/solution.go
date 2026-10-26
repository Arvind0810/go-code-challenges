package serializetree

// TreeNode is a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Serialize encodes the tree rooted at root into a string.
// Any format is allowed as long as Deserialize can reverse it.
func Serialize(root *TreeNode) string {
	// TODO: implement
	return ""
}

// Deserialize rebuilds the tree encoded by Serialize.
func Deserialize(data string) *TreeNode {
	// TODO: implement
	return nil
}
