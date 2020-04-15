package trees

//IsFull checks if the binary tree is full
func (n *Node) IsFull() bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	if n.Left != nil && n.Right == nil {
		return false
	}
	if n.Left == nil && n.Right != nil {
		return false
	}
	if n.Left != nil && n.Right != nil {
		return n.Left.IsFull() && n.Right.IsFull()
	}
	return false
}
