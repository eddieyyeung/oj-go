package inordertraversal

import (
	"container/list"
)

// TreeNode  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/explore/interview/card/top-interview-questions-medium/32/trees-and-graphs/85/
func inorderTraversal(root *TreeNode) []int {
	l := list.New()
	r := []int{}
	p := root
	for l.Len() != 0 || p != nil {
		if p != nil {
			l.PushBack(p)
			p = p.Left
		} else {
			e := l.Back()
			node, _ := e.Value.(*TreeNode)
			l.Remove(e)
			r = append(r, node.Val)
			p = node.Right
		}
	}
	return r
}
