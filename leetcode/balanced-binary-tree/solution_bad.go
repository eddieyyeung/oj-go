package balancedbinarytree

func isBalanced_bad(root *TreeNode) bool {
	var dfs func(node *TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		l_depth, l_isBalanced := dfs(node.Left)
		r_depth, r_isBalanced := dfs(node.Right)
		depth := max(l_depth, r_depth) + 1
		if !l_isBalanced || !r_isBalanced {
			return depth, false
		}
		if (l_depth-r_depth)*(l_depth-r_depth) <= 1 {
			return depth, true
		} else {
			return depth, false
		}
	}
	_, ans := dfs(root)
	return ans
}
