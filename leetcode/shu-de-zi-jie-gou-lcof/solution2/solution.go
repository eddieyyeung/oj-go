package solution2go

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A != nil && B != nil {
		return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
	}
	return false
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
