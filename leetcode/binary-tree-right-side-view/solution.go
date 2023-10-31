// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/description/
package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS 思路
// 遍历每一层，将最后一个节点的数值记录下来
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans []int
	for len(q) != 0 {
		ans = append(ans, q[len(q)-1].Val)
		var p []*TreeNode
		for _, v := range q {
			if v.Left != nil {
				p = append(p, v.Left)
			}
			if v.Right != nil {
				p = append(p, v.Right)
			}
		}
		q = p
	}
	return ans
}
