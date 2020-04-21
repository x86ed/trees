package trees

import "testing"

func TestIsSameTree(t *testing.T) {
	tree := newNode('a')
	tree.Left = newNode('b')
	tree.Right = newNode('d')
	tree.Left.Left = newNode('c')
	tree.Right.Right = newNode('e')

	sub := newNode('a')
	sub.Left = newNode('b')
	sub.Left.Left = newNode('c')
	sub.Right = newNode('d')

	got := tree.IsSameTree(tree)
	if got != true {
		t.Errorf("want tree.IsSameTree(tree) == true got false")
	}

	got = tree.IsSameTree(sub)
	if got != false {
		t.Errorf("want tree.IsSameTree(sub) == true got false")
	}
}
