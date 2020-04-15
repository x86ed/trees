package trees

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Stack []int

//IsEmpty  check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Top Returns the value at the top of the stack
func (s *Stack) Top() int {
	temp := []int(*s)
	return temp[len(temp)-1]
}

//Push a new value onto the stack
func (s *Stack) Push(i int) {
	*s = append(*s, i) // Simply append the new value to the end of the stack
}

//Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

//Empty tells if a stack is empty or not.
func (s *Stack) Empty() bool {
	if len([]int(*s)) == 0 {
		return true
	}
	return false
}

func map2slice(m map[int][]int) []int {
	out := []int{}
	var keys []int

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		out = append(out, m[v][0])
	}
	return out
}

func printInorder(n *Node) []int {
	if n == nil {
		return nil
	}
	out := []int{}
	if n != nil {
		out = append(out, printInorder(n.Left)...)
		out = append(out, n.Data)
		out = append(out, printInorder(n.Right)...)
	}
	return out
}

func printInOrderStr(n *Node) []string {
	if n == nil {
		return nil
	}
	out := []string{}
	if n != nil {
		out = append(out, printInOrderStr(n.Left)...)
		out = append(out, string(n.Data))
		out = append(out, printInOrderStr(n.Right)...)
	}
	return out
}

func printPreorder(n *Node) []int {
	if n == nil {
		return nil
	}
	out := []int{}
	if n != nil {
		out = append(out, n.Data)
		out = append(out, printPreorder(n.Left)...)
		out = append(out, printPreorder(n.Right)...)
	}
	return out
}

func printPreOrderStr(n *Node) []string {
	if n == nil {
		return nil
	}
	out := []string{}
	if n != nil {
		out = append(out, string(n.Data))
		out = append(out, printPreOrderStr(n.Left)...)
		out = append(out, printPreOrderStr(n.Right)...)
	}
	return out
}
