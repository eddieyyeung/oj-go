// 03. 二叉树的锯齿形层序遍历
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	isEven := false
	for len(q) != 0 {
		n := len(q)
		level := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if isEven {
				level[n-1-i] = node.Val
			} else {
				level[i] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, level)
		isEven = !isEven
	}
	return ans
}
