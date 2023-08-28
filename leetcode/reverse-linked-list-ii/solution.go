package reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	i := 0
	sentry := &ListNode{Next: head}
	p := sentry
	var prevRHead, tail *ListNode
	for p.Next != nil {
		i++
		if i < left {
			p = p.Next
		} else if i == left {
			prevRHead = p
			p = p.Next
			tail = p
		} else if i > left && i <= right {
			next := p.Next
			tail.Next = next.Next
			next.Next = prevRHead.Next
			prevRHead.Next = next
			p = tail
		} else {
			break
		}
	}
	return sentry.Next
}
