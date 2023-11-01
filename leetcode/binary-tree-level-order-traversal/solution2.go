package binarytreelevelordertraversal

// 两个队列
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans [][]int
	for len(q) != 0 {
		n := len(q)
		nq := []*TreeNode{}
		level := make([]int, n)
		for i, node := range q {
			level[i] = node.Val
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		ans = append(ans, level)
		q = nq
	}
	return ans
}
