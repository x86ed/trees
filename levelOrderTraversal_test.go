package trees

import (
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	root := newNode(20)
	root.Left = newNode(8)
	root.Right = newNode(22)
	root.Left.Left = newNode(5)
	root.Left.Right = newNode(3)
	root.Right.Left = newNode(4)
	root.Right.Right = newNode(25)
	root.Left.Right.Left = newNode(10)
	root.Left.Right.Right = newNode(14)
	got := root.LevelOrderTraversal()
	want := [][]int{[]int{20}, []int{8, 22}, []int{5, 3, 4, 25}, []int{10, 14}}
	for i, v := range got {
		for j, vv := range v {
			if vv != want[i][j] {
				t.Errorf("root.LevelOrderTraversal() = %v; want [][]int{[]int{20}, []int{8, 22}, []int{5, 3, 4, 25}, []int{10, 14}}", got)
			}
		}
	}
}
