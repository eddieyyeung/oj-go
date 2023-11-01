// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 一个队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans [][]int
	for len(q) != 0 {
		n := len(q)
		level := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			level[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
