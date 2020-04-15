package trees

//Leaves prints all leaves of a tree
func (n *Node) Leaves(a *[]int) []int {
	if n.Left == nil && n.Right == nil {
		*a = append(*a, n.Data)
	}
	if n.Left != nil {
		n.Left.Leaves(a)
	}
	if n.Right != nil {
		n.Right.Leaves(a)
	}
	return *a
}
