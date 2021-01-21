package tree

func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		Right: right,
		Left:  left,
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
	return true
}
