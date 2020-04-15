package trees

import "testing"

func TestMaxPathSum(t *testing.T) {
	root := newNode(10)
	root.Left = newNode(2)
	root.Right = newNode(10)
	root.Left.Left = newNode(20)
	root.Left.Right = newNode(1)
	root.Right.Right = newNode(-25)
	root.Right.Right.Left = newNode(3)
	root.Right.Right.Right = newNode(4)
	got := root.MaxPathSum()
	if got != 43 {
		t.Errorf("root.MaxPathSum() = %d; want 42", got)
	}
}
