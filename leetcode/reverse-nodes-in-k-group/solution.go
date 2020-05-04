// Package reverse_nodes_in_k_group ...
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	c := 0
	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	preTail := root
	p := root
	tail := root
	q := head
	for q != nil {
		if c%k == 0 {
			preTail = tail
			tail = q
			p = nil
		}
		tq := q.Next
		q.Next = p
		p = q
		preTail.Next = p
		q = tq
		c++
	}
	if c%k != 0 {
		q = preTail.Next
		p = nil
		for q != nil {
			tq := q.Next
			q.Next = p
			p = q
			preTail.Next = p
			q = tq
		}
	}
	return root.Next
}
