package count_good_nodes_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int

func goodNodes(root *TreeNode) int {
	count = 0

	max := root.Val
	FindMax(root, &max)
	return count
}

func FindMax(tn *TreeNode, max *int) {
	val := tn.Val
	tmp := *max
	if val > *max {
		*max = val
	}
	if tn.Right != nil {
		FindMax(tn.Right, max)
	}
	if tn.Left != nil {
		FindMax(tn.Left, max)
	}
	if val >= *max {
		count++
	}
	*max = tmp
}
