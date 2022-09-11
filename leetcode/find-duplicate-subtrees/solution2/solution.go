package solution2

// problem: https://leetcode.cn/problems/find-duplicate-subtrees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var records map[[3]int]int
var repeat map[int]*TreeNode
var idx int

func dfs(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	pk := [3]int{tn.Val, dfs(tn.Left), dfs(tn.Right)}
	if r, ok := records[pk]; ok {
		repeat[r] = tn
		return r
	} else {
		idx++
		records[pk] = idx
	}
	return idx
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	records = map[[3]int]int{}
	repeat = make(map[int]*TreeNode)
	idx = 0

	dfs(root)

	ans := make([]*TreeNode, 0, len(repeat))
	for _, v := range repeat {
		ans = append(ans, v)
	}
	return ans
}
