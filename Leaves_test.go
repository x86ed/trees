package trees

import (
	"testing"
)

func TestLeaves(t *testing.T) {
	root := newNode(20)
	root.Left = newNode(8)
	root.Right = newNode(22)
	root.Left.Left = newNode(5)
	root.Left.Right = newNode(3)
	root.Right.Left = newNode(4)
	root.Right.Right = newNode(25)
	root.Left.Right.Left = newNode(10)
	root.Left.Right.Right = newNode(14)
	buffer := []int{}
	got := root.Leaves(&buffer)
	want := []int{5, 10, 14, 4, 25}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("TestLeaves() = %v; want false", got)
		}
	}
}
