// Package intersection_of_two_linked_lists ...
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
package intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	p := headA
	for p != nil {
		m[p] = true
		p = p.Next
	}
	p = headB
	for p != nil {
		if m[p] {
			return p
		}
		p = p.Next
	}
	return nil
}
