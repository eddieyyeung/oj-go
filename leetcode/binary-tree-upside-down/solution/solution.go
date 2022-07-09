package solution

import "oj-go/utils"

var (
	logger = utils.NewLogger("binary-tree-upside-down")
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r *TreeNode

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	r = nil
	fn1(root)
	return r
}

func fn1(node *TreeNode) {
	if node == nil {
		return
	}
	left := node.Left
	if r == nil && left == nil {
		r = node
	}
	right := node.Right
	fn1(node.Left)
	fn1(node.Right)
	if left != nil {
		left.Right = node
		left.Left = right
	}
	node.Left = nil
	node.Right = nil
}
