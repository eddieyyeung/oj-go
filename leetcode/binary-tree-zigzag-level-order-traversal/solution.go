// Package binarytreezigzaglevelordertraversal ...
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
package binarytreezigzaglevelordertraversal

import (
	"container/list"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	r := new([][]int)
	l := list.New()
	if root == nil {
		return [][]int{}
	}
	l.PushBack(root)
	traversal(l, 1, r)
	return *r
}

func traversal(l *list.List, ltor int, r *[][]int) {
	newL := list.New()
	s := []int{}
	for e := l.Back(); e != nil; e = e.Prev() {
		tn, _ := e.Value.(*TreeNode)
		s = append(s, tn.Val)
		if ltor == 1 {
			if tn.Left != nil {
				newL.PushBack(tn.Left)
			}
			if tn.Right != nil {
				newL.PushBack(tn.Right)
			}
		}
		if ltor == -1 {
			if tn.Right != nil {
				newL.PushBack(tn.Right)
			}
			if tn.Left != nil {
				newL.PushBack(tn.Left)
			}
		}

	}
	*r = append(*r, s)
	if newL.Len() != 0 {
		traversal(newL, -ltor, r)
	}
}
