package trees

import "testing"

func TestMinDepth(t *testing.T) {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	got := root.MinDepth()
	if got != 2 {
		t.Errorf("root.MinDepth() = %d; want 2", got)
	}
}
