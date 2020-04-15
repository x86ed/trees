package trees

//TopNodes returns a sorted list from left to right ot the top layer of nodes when the tree is viewed from the top
func (n *Node) TopNodes() []int {
	if n == nil {
		return []int{}
	}
	b := &map[int][]int{}
	n.getTopNodes(b, 0, 0)
	return map2slice(*b)
}

func (n *Node) getTopNodes(bottom *map[int][]int, horiz, vert int) (left, right int) {
	t := *bottom
	if t[horiz] == nil || t[horiz][1] > vert {
		t[horiz] = []int{n.Data, vert}
	}
	bottom = &t
	vert++
	if n.Left != nil {
		lh := horiz - 1
		left, _ = n.Left.getTopNodes(bottom, lh, vert)
	}
	if n.Right != nil {
		rh := horiz + 1
		_, right = n.Right.getTopNodes(bottom, rh, vert)
	}
	return
}
