package trees

//Node object is the implementation of a binary tree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(i int) *Node {
	return &Node{
		Data: i,
	}
}

func (n *Node) notLeaf() bool {
	if n.Left != nil || n.Right != nil {
		return true
	}
	return false
}
