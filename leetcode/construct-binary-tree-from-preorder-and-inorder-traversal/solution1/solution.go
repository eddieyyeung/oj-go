// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description
package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	len_pre, len_in := len(preorder), len(inorder)
	var dfs func(left_pre, right_pre, left_in, right_in int) *TreeNode

	dfs = func(left_pre, right_pre, left_in, right_in int) *TreeNode {
		if left_pre >= right_pre {
			return nil
		}
		mid := preorder[left_pre]
		mid_node := TreeNode{
			Val: mid,
		}
		idx := 0
		for i := left_in; i < right_in; i++ {
			if inorder[i] == mid {
				idx = i
			}
		}
		mid_node.Left = dfs(left_pre+1, idx-left_in+left_pre+1, left_in, idx)
		mid_node.Right = dfs(idx-left_in+left_pre+1, right_pre, idx+1, right_in)
		return &mid_node
	}
	return dfs(0, len_pre, 0, len_in)
}
