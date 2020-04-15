package trees

import "testing"

func TestCanBeBinarySearchTree(t *testing.T) {
	pre1 := []int{40, 30, 35, 80, 100}
	got := CanBeBinarySearchTree(pre1)
	if got != true {
		t.Errorf("CanBeBinarySearchTree() = %v; want true", got)
	}
	pre2 := []int{40, 30, 35, 20, 80, 100}
	got = CanBeBinarySearchTree(pre2)
	if got != false {
		t.Errorf("CanBeBinarySearchTree() = %v; want false", got)
	}
}
