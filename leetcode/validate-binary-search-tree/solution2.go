package validatebinarysearchtree

import "math"

// 中序遍历：先左节点，当前节点，后右节点
// 由于遍历完之后数值是递增的，因此可以用额外的变量来判断是否符合递增规律
func isValidBST2(root *TreeNode) bool {
	pre := -math.MaxInt
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}
