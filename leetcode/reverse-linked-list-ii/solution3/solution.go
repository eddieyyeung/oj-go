package solution3

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	p0 := sentry
	for i := 1; i <= left-1; i++ {
		p0 = p0.Next
	}
	cur := p0.Next
	var pre *ListNode
	for i := left; i <= right; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	p0.Next.Next = cur
	p0.Next = pre
	return sentry.Next
}
