// Package lowest_common_ancestor_of_a_binary_tree ...
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
package lowest_common_ancestor_of_a_binary_tree

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var dfs func(node, p, q *TreeNode) int
	dfs = func(node, p, q *TreeNode) int {
		if node == nil {
			return 0
		}
		l_cnt := dfs(node.Left, p, q)
		r_cnt := dfs(node.Right, p, q)
		cnt := 0
		if node == p || node == q {
			cnt++
		}
		cnt += l_cnt + r_cnt
		if ans == nil && cnt == 2 {
			ans = node
		}
		return cnt
	}
	dfs(root, p, q)
	return ans
}
