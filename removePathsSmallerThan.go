package trees

import "fmt"

//RemovePathsSmallerThan removes nodes part of paths smaller than x
func (n *Node) RemovePathsSmallerThan(x int) {
	d := 1
	n.removePathsUtil(d, x, nil, false)
}

func (n *Node) removePathsUtil(d, x int, p *Node, l bool) {
	if n == nil {
		return
	}
	if n.Left != nil {
		next := d + 1
		n.Left.removePathsUtil(next, x, n, true)
	}
	if n.Right != nil {
		next := d + 1
		n.Right.removePathsUtil(next, x, n, false)
	}
	if n.Left == nil && n.Right == nil && d < x && p != nil {
		var dead *Node
		fmt.Println(n, d, x, p)
		if l {
			p.Left = dead
		} else {
			p.Right = dead
		}
	}
}
