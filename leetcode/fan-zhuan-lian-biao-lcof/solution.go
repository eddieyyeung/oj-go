package fanzhuanlianbiaolcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainningPlan(head *ListNode) *ListNode {
	var root *ListNode
	p := head
	for p != nil {
		next := p.Next
		p.Next = root
		root = p
		p = next
	}
	return root
}
