package solution1

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	v := -101
	sentry := &ListNode{Next: head}
	p0, p1 := sentry, head
	for p1 != nil {
		nxt := p1.Next
		if p1.Val != v {
			p0 = p0.Next
			v = p1.Val
		} else {
			p0.Next = p0.Next.Next
		}
		p1 = nxt
	}
	return sentry.Next
}
