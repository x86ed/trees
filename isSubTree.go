package trees

import "strings"

//IsSubTree finds out if the tree entered is a subtree of the main tree
func (n *Node) IsSubTree(s *Node) bool {
	if n == nil || s == nil {
		return false
	}

	nio := printInOrderStr(n)
	sio := printInOrderStr(s)
	nios := strings.Join(nio, " ")
	sios := strings.Join(sio, " ")
	if strings.Contains(nios, sios) {
		return true
	}

	nio = printPreOrderStr(n)
	sio = printPreOrderStr(s)
	nios = strings.Join(nio, " ")
	sios = strings.Join(sio, " ")
	if strings.Contains(nios, sios) {
		return true
	}
	return false
}
