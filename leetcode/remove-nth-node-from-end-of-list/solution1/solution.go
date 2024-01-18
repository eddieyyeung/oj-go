package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	p, q := sentry, sentry
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	q.Next = q.Next.Next
	return sentry.Next
}
