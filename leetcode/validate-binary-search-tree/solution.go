package validatebinarysearchtree

import "math"

// 前序遍历：先判断，再递归
// 1. 判断是否在区间范围内
// 2. 将区间范围传入到左右节点，左：(left, val)，右：(val, right)
// 3. 初始化区间为：(-inf, inf)
func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left int, right int) bool
	dfs = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}
		val := node.Val
		return val > left && val < right && dfs(node.Left, left, val) && dfs(node.Right, val, right)
	}
	return dfs(root, -math.MaxInt, math.MaxInt)
}
