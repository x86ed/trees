package trees

//Reverse reverses the order of a binary tree
func (n *Node) Reverse() {
	n.reverseSwap(n.Left, n.Right, 0)
}

func (n *Node) reverseSwap(n1, n2 *Node, l int) {
	if n.Left == nil || n.Right == nil {
		return
	}

	if l%2 == 0 {
		n.Left, n.Right = n2, n1
	}
	next := l + 1
	n1.reverseSwap(n1.Left, n2.Right, next)
	n2.reverseSwap(n1.Right, n2.Left, next)
}
