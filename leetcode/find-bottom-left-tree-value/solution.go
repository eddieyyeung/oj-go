// 513. 找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value/description/
package findbottomlefttreevalue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	var ans int
	for len(q) != 0 {
		n := len(q)
		ans = q[0].Val
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return ans
}
