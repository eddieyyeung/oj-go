// Package solution2 ...
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
package solution2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := map[*ListNode]bool{}
	p := head
	m[p] = true
	for {
		if p.Next == nil {
			return nil
		}
		p = p.Next
		if m[p] {
			return p
		}
		m[p] = true
	}
}
