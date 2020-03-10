// Package diameter_of_binary_tree ...
// https://leetcode-cn.com/problems/diameter-of-binary-tree/
package diameter_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	md := 0
	findDepth(root, &md)
	return md
}

func findDepth(tn *TreeNode, md *int) int {
	ld := 0
	rd := 0
	if tn.Left != nil {
		ld = findDepth(tn.Left, md) + 1
	}
	if tn.Right != nil {
		rd = findDepth(tn.Right, md) + 1
	}
	m := max(ld, rd)
	if sum := ld + rd; sum > *md {
		*md = sum
	}
	return m
}

func max(a ...int) int {
	max := math.MinInt64
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
