package congshangdaoxiadayinerchashulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	q := []*TreeNode{root}
	var ans []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			ans = append(ans, node.Val)
			q = append(q, node.Left, node.Right)
		}
	}
	return ans
}
