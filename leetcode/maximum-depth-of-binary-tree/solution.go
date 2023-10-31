// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用递归的思想：
// 1. 把原问题转换为子问题
// 2. 边界条件
func maxDepth(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l_depth := dfs(node.Left)
		r_depth := dfs(node.Right)
		return max(l_depth, r_depth) + 1
	}
	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
