package trees

//BottomNodes prints all the nodes of the tree unobstructed if the tree is viewed from below
func (n *Node) BottomNodes() []int {
	if n == nil {
		return []int{}
	}
	b := &map[int][]int{}
	n.getBottomNodes(b, 0, 0)
	return map2slice(*b)
}

func (n *Node) getBottomNodes(bottom *map[int][]int, horiz, vert int) (left, right int) {
	t := *bottom
	if t[horiz] == nil || t[horiz][1] < vert {
		t[horiz] = []int{n.Data, vert}
	}
	bottom = &t
	vert++
	if n.Left != nil {
		lh := horiz - 1
		left, _ = n.Left.getBottomNodes(bottom, lh, vert)
	}
	if n.Right != nil {
		rh := horiz + 1
		_, right = n.Right.getBottomNodes(bottom, rh, vert)
	}
	return
}
