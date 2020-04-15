package trees

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := newNode(20)
	root.Left = newNode(8)
	root.Right = newNode(22)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(12)
	root.Left.Right.Left = newNode(10)
	root.Left.Right.Right = newNode(14)

	n1 := 10
	n2 := 14
	got := root.LowestCommonAncestor(n1, n2)
	want := 12
	if got != want {
		t.Errorf("want root.LowestCommonAncestor(%d,%d) == %d got %d", n1, n2, want, got)
	}

	n1 = 14
	n2 = 8
	got = root.LowestCommonAncestor(n1, n2)
	want = 8
	if got != want {
		t.Errorf("want root.LowestCommonAncestor(%d,%d) == %d got %d", n1, n2, want, got)
	}

	n1 = 10
	n2 = 22
	got = root.LowestCommonAncestor(n1, n2)
	want = 20
	if got != want {
		t.Errorf("want root.LowestCommonAncestor(%d,%d) == %d got %d", n1, n2, want, got)
	}
}
