package trees

//IsSameTree checks if the tree passed in is identical to the one being tested against
func (n *Node) IsSameTree(o *Node) bool {
	if n == nil && o == nil {
		return true
	}
	if n == nil || o == nil {
		return false
	}
	if n.Data != o.Data {
		return false
	}
	return n.Left.IsSameTree(o.Left) && n.Right.IsSameTree(o.Right)
}
