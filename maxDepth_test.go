package trees

import "testing"

func TestMaxDepth(t *testing.T) {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Right.Left = newNode(5)
	root.Right.Right = newNode(6)
	root.Right.Right.Left = newNode(8)
	root.Right.Left.Right = newNode(7)

	got := root.MaxDepth()
	if got != 4 {
		t.Errorf("root.MaxDepth() = %d want %d", got, 4)
	}
}
