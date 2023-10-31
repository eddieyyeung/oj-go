// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 类似求二叉树的高度，计算出左右节点的高度，并比较高度差的绝对值是否大于 1
// 2. 利用树的高度始终 >= 0 的特性，若判断不平衡，则返回 -1
func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l_depth := dfs(node.Left)
		if l_depth == -1 {
			return -1
		}
		r_depth := dfs(node.Right)
		if r_depth == -1 {
			return -1
		}
		if abs(l_depth, r_depth) > 1 {
			return -1
		}
		return max(l_depth, r_depth) + 1
	}
	return dfs(root) != -1
}

func abs(a, b int) int {
	if b > a {
		a, b = b, a
	}
	return a - b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
