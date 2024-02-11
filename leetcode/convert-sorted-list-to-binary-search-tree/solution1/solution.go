// 109. 有序链表转换二叉搜索树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
package solution

// 解法：
// 1. 找中位数
// 2. 分治

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func findMidListNode(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast.Next != right && fast.Next.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	midListNode := findMidListNode(left, right)
	tnode := &TreeNode{
		Val: midListNode.Val,
	}
	tnode.Left = buildTree(left, midListNode)
	tnode.Right = buildTree(midListNode.Next, right)
	return tnode
}
