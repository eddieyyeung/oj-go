// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/description/
package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 类似 100. 相同的树 https://leetcode.cn/problems/same-tree/description/
func isSymmetric(root *TreeNode) bool {
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}
	return dfs(root.Left, root.Right)
}
