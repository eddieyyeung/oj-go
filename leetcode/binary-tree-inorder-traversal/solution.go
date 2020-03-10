// Package binarytreeinordertraversal ...
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
package binarytreeinordertraversal

import (
	"container/list"
)

// TreeNode  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func preorderTraversal(root *TreeNode) []int {
	l := list.New()
	r := []int{}
	p := root
	for l.Len() != 0 || p != nil {
		if p != nil {
			r = append(r, p.Val)
			l.PushBack(p)
			p = p.Left
		} else {
			e := l.Back()
			node, _ := e.Value.(*TreeNode)
			l.Remove(e)
			p = node.Right
		}
	}
	return r
}
