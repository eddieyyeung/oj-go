// 111. 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l_depth := dfs(node.Left)
		r_depth := dfs(node.Right)
		var rst = 1
		if l_depth != 0 && r_depth != 0 {
			rst = min(l_depth, r_depth) + 1
		} else if l_depth != 0 {
			rst = l_depth + 1
		} else if r_depth != 0 {
			rst = r_depth + 1
		}
		return rst
	}
	return dfs(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
