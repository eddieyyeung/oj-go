// Package lowest_common_ancestor_of_a_binary_tree ...
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	l := root
	_, _, r := dfs(l, p, q)
	return r
}

func dfs(l, p, q *TreeNode) (bool, bool, *TreeNode) {
	var lFoundP, lFoundQ, rFoundP, rFoundQ bool
	var fTreeNode, rTreeNode *TreeNode
	if l.Left != nil {
		lFoundP, lFoundQ, fTreeNode = dfs(l.Left, p, q)
	}
	if l.Right != nil {
		rFoundP, rFoundQ, rTreeNode = dfs(l.Right, p, q)
	}
	foundP := false
	foundQ := false
	var r *TreeNode
	if lFoundP || rFoundP {
		foundP = true
	}
	if lFoundQ || rFoundQ {
		foundQ = true
	}
	if fTreeNode != nil {
		r = fTreeNode
	}
	if rTreeNode != nil {
		r = rTreeNode
	}
	if l == p {
		foundP = true
	}
	if l == q {
		foundQ = true
	}
	if foundP && foundQ && r == nil {
		r = l
	}
	return foundP, foundQ, r
}
