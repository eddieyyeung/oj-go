// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
package constructbinarytreefrominorderandpostordertraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var dfs func(inorder []int, postorder []int) *TreeNode
	dfs = func(inorder, postorder []int) *TreeNode {
		if len(inorder) == 0 {
			return nil
		}

		last_post := postorder[len(postorder)-1]
		node := &TreeNode{
			Val: last_post,
		}
		var idx int = -1
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == last_post {
				idx = i
				break
			}
		}
		node.Left = dfs(inorder[:idx], postorder[:idx])
		node.Right = dfs(inorder[idx+1:], postorder[idx:len(postorder)-1])
		return node
	}
	return dfs(inorder, postorder)
}
