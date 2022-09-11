package trim_a_binary_search_tree

// problem: https://leetcode.cn/problems/trim-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	sentinel := &TreeNode{
		Val:   0,
		Left:  root,
		Right: nil,
	}
	dfs(sentinel, low, high)
	return sentinel.Left
}

func dfs(tn *TreeNode, low int, high int) {
	if tn == nil {
		return
	}

	if tn.Left != nil {
		if tn.Left.Val < low {
			tn.Left = tn.Left.Right
			dfs(tn, low, high)
		}
		dfs(tn.Left, low, high)
	}
	if tn.Right != nil {
		if tn.Right.Val > high {
			tn.Right = tn.Right.Left
			dfs(tn, low, high)
		}
		dfs(tn.Right, low, high)
	}
}
