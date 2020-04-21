package trees

//LevelOrderTraversal creates a nested list of the contents of eaxh level
func (n *Node) LevelOrderTraversal() [][]int {
	if n == nil {
		return nil
	}
	out := [][]int{
		[]int{n.Data},
	}
	out = n.Left.levelOrderTraversalUtil(out, 1)
	out = n.Right.levelOrderTraversalUtil(out, 1)
	return out
}

func (n *Node) levelOrderTraversalUtil(arr [][]int, h int) [][]int {
	if n == nil {
		return arr
	}
	for len(arr) < h+1 {
		arr = append(arr, []int{})
	}
	arr[h] = append(arr[h], n.Data)
	if n.Left != nil {
		arr = n.Left.levelOrderTraversalUtil(arr, h+1)
	}
	if n.Right != nil {
		arr = n.Right.levelOrderTraversalUtil(arr, h+1)
	}
	return arr
}
