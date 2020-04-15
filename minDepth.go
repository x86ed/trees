package trees

//MinDepth gives the minimum depth of a binary tree
func (n *Node) MinDepth() int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	if n.Right == nil {
		return n.Left.MinDepth() + 1
	}
	if n.Left == nil {
		return n.Right.MinDepth() + 1
	}
	return min(n.Left.MinDepth(), n.Right.MinDepth()) + 1
}
