package trees

//CanBeBinarySearchTree takes an array of integers and returns if it's a traversal of a binary search tree.
func CanBeBinarySearchTree(a []int) bool {
	s := Stack{}

	root := MinInt
	for _, v := range a {
		if v < root {
			return false
		}
		for !s.Empty() && s.Top() < v {
			root = s.Top()
			s.Pop()
		}
		s.Push(v)
	}
	return true
}
