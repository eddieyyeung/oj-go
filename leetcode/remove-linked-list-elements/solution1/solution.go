// 203. 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/description/
package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sentry := &ListNode{Next: head}
	p := sentry
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return sentry.Next
}
