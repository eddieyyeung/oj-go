package duichengdeerchashulcof

// 剑指 Offer 28. 对称的二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
