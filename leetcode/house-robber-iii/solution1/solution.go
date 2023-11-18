// 337. 打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/description/
package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		l_rob, l_norob := dfs(node.Left)
		r_rob, r_norob := dfs(node.Right)
		rob := l_norob + r_norob + node.Val
		norob := max(l_rob, l_norob) + max(r_rob, r_norob)
		return rob, norob
	}
	rob, norob := dfs(root)
	return max(rob, norob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
