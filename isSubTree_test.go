package trees

import "testing"

func TestIsSubTree(t *testing.T) {
	tree := newNode('a')
	tree.Left = newNode('b')
	tree.Right = newNode('d')
	tree.Left.Left = newNode('c')
	tree.Right.Right = newNode('e')

	sub := newNode('a')
	sub.Left = newNode('b')
	sub.Left.Left = newNode('c')
	sub.Right = newNode('d')

	got := tree.IsSubTree(sub)
	if got != true {
		t.Errorf("want tree.IsSubTree(sub) == true got false")
	}
}
