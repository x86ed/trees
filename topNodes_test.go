package trees

import "testing"

func TestTopNodes(t *testing.T) {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Right = newNode(4)
	root.Left.Right.Right = newNode(5)
	root.Left.Right.Right.Right = newNode(6)
	got := root.TopNodes()
	want := []int{2, 1, 3, 6}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("TestTopNodes() = %v; want [2 1 3 6]", got)
		}
	}
}
