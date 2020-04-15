package trees

//LowestCommonAncestor prints the closest ancestor between two values in a binary search tree (values must be present in the tree)
func (n *Node) LowestCommonAncestor(v, m int) int {
	for n != nil {
		if n.Data > v && n.Data > m {
			n = n.Left
			// If both n1 and n2 are greater than root, then LCA lies in right
		} else if n.Data < v && n.Data < m {
			n = n.Right
		} else {
			break
		}
	}
	return n.Data
}
