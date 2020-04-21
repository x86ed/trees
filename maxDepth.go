package trees

//MaxDepth returns the maximum distance from a leaf to the root of a given tree
func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	}

	leftDepth := n.Left.MaxDepth()
	rightDepth := n.Right.MaxDepth()

	// Choose the larger one and add the current depth to it.
	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
