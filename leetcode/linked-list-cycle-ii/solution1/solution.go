// Package solution1 ...
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
package solution1

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	q := head
	for {
		if p.Next == nil || q.Next == nil || q.Next.Next == nil {
			return nil
		}
		p = p.Next
		q = q.Next.Next
		if p == q {
			break
		}
	}
	p = head
	for {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}
}
