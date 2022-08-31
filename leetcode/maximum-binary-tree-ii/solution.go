package maximum_binary_tree_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	ntn := TreeNode{Val: val}
	if root == nil {
		return &ntn
	}
	if ntn.Val > root.Val {
		ntn.Left = root
		return &ntn
	}
	RecurInsertIntoTree(root, &ntn)
	return root
}

func RecurInsertIntoTree(tn *TreeNode, ntn *TreeNode) {
	if tn.Right != nil {
		if ntn.Val > tn.Right.Val {
			ntn.Left = tn.Right
			tn.Right = ntn
			return
		} else {
			RecurInsertIntoTree(tn.Right, ntn)
		}
	} else {
		tn.Right = ntn
		return
	}
}
