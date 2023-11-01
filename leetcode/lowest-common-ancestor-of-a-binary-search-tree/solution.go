// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 如果 p 和 q 都在左子树，则返回递归左子树的结果
// 2. 如果 p 和 q 都在右子树，则返回递归右子树的结果
// 3. 如果 p 和 q 分别在左右子树，则返回当前节点
// 4. 如果当前节点是 p，则返回当前节点
// 5. 如果当前节点是 q，则返回当前节点
// 6. 题目保证 p, q 为不同节点且存在于给定的二叉搜索树中，所以不需要判断空节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(node, p, q *TreeNode) *TreeNode
	dfs = func(node, p, q *TreeNode) *TreeNode {
		if p.Val < node.Val && q.Val < node.Val {
			return dfs(node.Left, p, q)
		}
		if p.Val > node.Val && q.Val > node.Val {
			return dfs(node.Right, q, p)
		}
		return node
	}

	return dfs(root, p, q)
}
