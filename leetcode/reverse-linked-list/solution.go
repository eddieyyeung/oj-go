// Package reverse_linked_list ...
// https://leetcode-cn.com/problems/reverse-linked-list/
package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p, q *ListNode
	p = head
	for p != nil {
		t := p
		p = p.Next
		t.Next = q
		q = t
	}
	return q
}
