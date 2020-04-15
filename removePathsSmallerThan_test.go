package trees

import "testing"

func TestRemovePathsSmallerThan(t *testing.T){
    k := 4
    root := newNode(1) 
    root.Left = newNode(2) 
    root.Right = newNode(3) 
    root.Left.Left = newNode(4) 
    root.Left.Right = newNode(5) 
    root.Left.Left.Left = newNode(7) 
    root.Right.Right = newNode(6) 
    root.Right.Right.Left = newNode(8) 
    got := printInorder(root) 
    want := []int{7,4, 2, 5, 1, 3, 8, 6}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("TestLeaves() = %v; want [7 4 2 5 1 3 8 6]", got)
		}
	}
    root.RemovePathsSmallerThan(k)
    got = printInorder(root) 
    want = []int{7,4,2,1,3,8,6}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("TestLeaves() = %v; want [7 4 2 1 3 8 6]", got)
		}
	}
}