package longest_univalue_path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var longest int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	longest = 0
	c := 1
	recurFindLongestUnivaluePath(root, root.Val, &c)
	return longest - 1
}

func recurFindLongestUnivaluePath(tn *TreeNode, prev int, c *int) {
	l := *c
	r := *c
	if tn.Left != nil {
		recurFindLongestUnivaluePath(tn.Left, tn.Val, &l)
	}
	if tn.Right != nil {
		recurFindLongestUnivaluePath(tn.Right, tn.Val, &r)
	}
	if tn.Val != prev {
		newC := 1
		c = &newC
	} else {
		*c++
	}
	if result := l + r - 1; result > longest {
		longest = result
	}
	if l > r {
		*c = l + 1
	} else {
		*c = r + 1
	}
}
