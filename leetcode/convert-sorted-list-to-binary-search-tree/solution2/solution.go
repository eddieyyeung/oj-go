// 109. 有序链表转换二叉搜索树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
package solution

// 解法：中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var gnode *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	length := getListNodeLength(head)
	gnode = head
	return buildTree(0, length-1)
}

func getListNodeLength(lnode *ListNode) int {
	var cnt int
	for lnode != nil {
		cnt++
		lnode = lnode.Next
	}
	return cnt
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	m := (left + right) / 2
	lnode := buildTree(left, m-1)
	node := &TreeNode{
		Val: gnode.Val,
	}
	gnode = gnode.Next
	rnode := buildTree(m+1, right)
	node.Left = lnode
	node.Right = rnode
	return node
}
