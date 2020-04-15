package trees

//MaxPathSum represents the maximum sum of any path in the given tree.
func (n *Node) MaxPathSum() int {
	res := MinInt
	result := &res
	out := maxPathSumUtil(n, result)
	return out
}

func maxPathSumUtil(n *Node, res *int) int {
	if n == nil {
		return 0
	}
	l := maxPathSumUtil(n.Left, res)
	r := maxPathSumUtil(n.Right, res)

	maxV := max(max(l, r)+n.Data, n.Data)

	maxR := max(maxV, l+r+n.Data)

	raw := max(*res, maxR)
	res = &raw

	return raw
}
