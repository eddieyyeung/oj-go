// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
package solution

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历二叉树，在计算最大链和的同时，顺带更新路径和的最大值
// 在当前节点的最大路径和 = 左子树最大链和 + 右子树最大链和 + 当前节点
// 如果这个值是负数，则返回 0
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxPathSum(root *TreeNode) int {
	var ans int = -math.MaxInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l_max := dfs(node.Left)
		r_max := dfs(node.Right)
		ans = max(ans, l_max+r_max+node.Val)
		rst := max(max(l_max+node.Val, r_max+node.Val), 0)
		return rst
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
