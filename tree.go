package tree

func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func IsBinarySearchTree(node *Node) bool {
	if node == nil {
		return true
	}
	if !(node.Left.Val < node.Val &&
		node.Right.Val > node.Val) {
		return false
	}
	return true
}
