package binarytreezigzaglevelordertraversal

func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	isEven := false
	for len(q) != 0 {
		nq := []*TreeNode{}
		level := make([]int, len(q))
		for i, node := range q {
			if isEven {
				level[len(q)-1-i] = node.Val
			} else {
				level[i] = node.Val
			}
			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
		q = nq
		ans = append(ans, level)
		isEven = !isEven
	}
	return ans
}
