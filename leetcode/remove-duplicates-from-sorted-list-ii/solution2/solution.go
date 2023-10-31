package solution2

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentry := &ListNode{Next: head}
	p0, p1 := sentry, head
	cnt := 0
	for p1.Next != nil {
		nxt := p1.Next
		if p1.Val == p1.Next.Val {
			p1 = nxt
			cnt++
			continue
		}
		if cnt != 0 {
			p0.Next = nxt
		} else {
			p0 = p1
		}
		cnt = 0
		p1 = nxt
	}
	if cnt != 0 {
		p0.Next = p1.Next
	}
	return sentry.Next
}
