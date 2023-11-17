// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/
package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历二叉树，在计算最长链的同时，顺带把直径算出来。
// 在当前节点的直径长度 = 左子树的最长链 + 右子树的最长链 + 2
// 注意：这里求的是链长，即边长，而不是深度
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		l_len := dfs(node.Left)
		r_len := dfs(node.Right)
		ans = max(ans, l_len+r_len+2)
		return max(l_len, r_len) + 1
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
