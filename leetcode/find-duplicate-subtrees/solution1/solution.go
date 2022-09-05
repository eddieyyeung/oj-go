package solution1

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NullVal = ""

var m map[string]int

var tns []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	m = make(map[string]int)
	tns = []*TreeNode{}
	dfs(root)
	return tns
}

func dfs(tn *TreeNode) string {
	if tn == nil {
		return NullVal
	}
	serial := strconv.Itoa(tn.Val) + "," + dfs(tn.Left) + "," + dfs(tn.Right)

	if count, ok := m[serial]; ok {
		if count == 1 {
			tns = append(tns, tn)
		}
	}
	m[serial]++
	return serial
}
