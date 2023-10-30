package solution2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	p0 := sentry
	cur := sentry.Next
	cnt := 1
	var pre *ListNode
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
		if cnt%k == 0 {
			p0.Next.Next = cur
			p0nxt := p0.Next
			p0.Next = pre
			p0 = p0nxt
			pre = nil
		}
		cnt++
	}
	if pre != nil {
		p1 := &ListNode{
			Next: pre,
		}
		cur = pre
		pre = nil
		for cur != nil {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		p1.Next.Next = cur
		p1.Next = pre
		p0.Next = p1.Next
	}

	return sentry.Next
}
