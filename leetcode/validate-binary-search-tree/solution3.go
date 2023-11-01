package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历：先递归，后判断
// 1. 递归后，分别返回左节点和右节点的区间范围
// 2. 判断当前节点是否在区间范围内，否则返回 [-inf, inf]
// 3. 遇到空节点返回 [inf, -inf] 保证一定满足条件
func isValidBST3(root *TreeNode) bool {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, -math.MaxInt
		}
		l_min, l_max := dfs(node.Left)
		r_min, r_max := dfs(node.Right)
		c_val := node.Val
		if c_val <= l_max || c_val >= r_min {
			return -math.MaxInt, math.MaxInt
		}
		return min(c_val, l_min), max(c_val, r_max)
	}
	a, _ := dfs(root)
	return a != -math.MaxInt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
