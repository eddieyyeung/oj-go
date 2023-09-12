package shudezijiegoulcof

import "container/list"

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
	l := list.New()
	l.PushBack(A)
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		v := e.Value.(*TreeNode)
		if v.Val == B.Val {
			if isSame(v, B) {
				return true
			}
		}
		if v.Left != nil {
			l.PushBack(v.Left)
		}
		if v.Right != nil {
			l.PushBack(v.Right)
		}
	}
	return false
}
