package utils

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateNode(arr []*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   *arr[0],
		Left:  nil,
		Right: nil,
	}
	l := list.New()
	l.PushBack(root)
	for i := 1; i < len(arr); i++ {
		e := l.Front()
		p := e.Value.(*TreeNode)
		var left, right *TreeNode
		if arr[i] != nil {
			left = &TreeNode{
				Val:   *arr[i],
				Left:  nil,
				Right: nil,
			}
			l.PushBack(left)
		}
		if i+1 < len(arr) {
			i = i + 1
			if arr[i] != nil {
				right = &TreeNode{
					Val:   *arr[i],
					Left:  nil,
					Right: nil,
				}
				l.PushBack(right)
			}
		}
		p.Left = left
		p.Right = right
		l.Remove(e)
	}
	return root
}
