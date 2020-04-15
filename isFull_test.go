package trees

import "testing"

func TestIsFull(t *testing.T) {
	root := newNode(10)
	root.Left = newNode(20)
	root.Right = newNode(30)

	root.Left.Right = newNode(40)
	root.Left.Left = newNode(50)
	root.Right.Left = newNode(60)
	root.Right.Right = newNode(70)

	root.Left.Left.Left = newNode(80)
	root.Left.Left.Right = newNode(90)
	root.Left.Right.Left = newNode(80)
	root.Left.Right.Right = newNode(90)
	root.Right.Left.Left = newNode(80)
	root.Right.Left.Right = newNode(90)
	root.Right.Right.Left = newNode(80)
	root.Right.Right.Right = newNode(90)
	got := root.IsFull()
	if got == false {
		t.Errorf("root.IsFull() == %v; want true", got)
	}
	root = newNode(10)
	root.Left = newNode(20)
	root.Right = newNode(30)

	root.Left.Right = newNode(40)
	root.Left.Left = newNode(50)
	root.Right.Left = newNode(60)
	root.Right.Right = newNode(70)

	root.Left.Left.Left = newNode(80)
	root.Left.Left.Right = newNode(90)
	root.Left.Right.Left = newNode(80)
	root.Left.Right.Right = newNode(90)
	root.Right.Left.Right = newNode(90)
	root.Right.Right.Left = newNode(80)
	root.Right.Right.Right = newNode(90)
	got = root.IsFull()
	if got == true {
		t.Errorf("root.IsFull() == %v; want false", got)
	}
}
