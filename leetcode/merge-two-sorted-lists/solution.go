// Package merge_two_sorted_lists ...
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	root := &ListNode{}
	t := root
	for p != nil && q != nil {
		if p.Val < q.Val {
			t.Next = p
			p = p.Next
		} else {
			t.Next = q
			q = q.Next
		}
		t = t.Next
	}
	for p != nil {
		t.Next = p
		t = t.Next
		p = p.Next
	}
	for q != nil {
		t.Next = q
		t = t.Next
		q = q.Next
	}
	return root.Next
}
