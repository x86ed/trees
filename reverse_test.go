package trees

import "testing"

func TestReverse(t *testing.T) {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(7)
	root.Left.Left.Left = newNode(8)
	root.Left.Left.Right = newNode(9)
	root.Left.Right.Left = newNode(10)
	root.Left.Right.Right = newNode(11)
	root.Right.Left.Left = newNode(12)
	root.Right.Left.Right = newNode(13)
	root.Right.Right.Left = newNode(14)
	root.Right.Right.Right = newNode(15)
	got := printPreorder(root)
	want := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("root.Reverse() = %v; want [1 2 4 8 9 5 10 11 3 6 12 13 7 14 15]", got)
		}
	}
	root.Reverse()
	got = printPreorder(root)
	want = []int{1, 3, 6, 12, 10, 7, 14, 8, 2, 4, 15, 8, 5, 13, 10}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("root.Reverse() = %v; want [1 3 6 12 10 7 14 8 2 4 15 8 5 13 10]", got)
		}
	}
}
