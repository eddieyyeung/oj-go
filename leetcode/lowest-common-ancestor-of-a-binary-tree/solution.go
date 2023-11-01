// 236. 二叉树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 如果当前节点为空节点，返回空节点
// 2. 如果当前节点为 p，返回 p
// 3. 如果当前节点为 q，返回 q
// 1～3 => 返回当前节点
// 4. 如果左右树都找到节点，返回当前节点
// 5. 如果只有左树找到节点，返回左树结果
// 6. 如果只有右树找到节点，返回右树结果
// 7. 如果左右树都找不到节点，返回空节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(node, p, q *TreeNode) *TreeNode
	dfs = func(node, p, q *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}
		l_node := dfs(node.Left, p, q)
		r_node := dfs(node.Right, p, q)
		if l_node != nil && r_node != nil {
			return node
		}
		if l_node != nil {
			return l_node
		}
		return r_node
	}

	return dfs(root, p, q)
}
