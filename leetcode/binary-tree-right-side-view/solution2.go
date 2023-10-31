package binarytreerightsideview

// 递归思路 DFS
// 1. 先遍历右指针，再遍历左指针
// 2. 判断每次递归的深度与 len(ans) 是否相等，若相等，则记录当前节点的数值
func rightSideView2(root *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}
