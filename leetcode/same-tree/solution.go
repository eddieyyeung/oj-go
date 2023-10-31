// 100. 相同的树
// https://leetcode.cn/problems/same-tree/description/
package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		// 情况 1: 如果 p q 都为 nil，则 返回 true
		// 情况 2: 如果 p q 其中一个为 nil，另一个不为 nil，则返回 false
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
	}
	return dfs(p, q)
}
