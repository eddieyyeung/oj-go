package shanchulianbiaodejiedianlcof

func deleteNode(head *ListNode, val int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}

	p := sentry

	for p.Next != nil {
		prev := p
		p = p.Next
		if p.Val == val {
			prev.Next = p.Next
			p.Next = nil
			break
		}
	}
	return sentry.Next
}
