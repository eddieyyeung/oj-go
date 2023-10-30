package solution2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	p0 := sentry
	curr := head
	var revHead, tail *ListNode
	cnt := 0
	for curr != nil {
		cnt++
		if cnt < left {
			p0 = curr
			curr = curr.Next
			continue
		}
		if cnt == left {
			tail = curr
		}
		if cnt >= left && cnt <= right {
			next := curr.Next
			curr.Next = revHead
			revHead = curr
			curr = next
		}
		if cnt > right {
			break
		}
	}

	tail.Next = curr
	p0.Next = revHead
	return sentry.Next
}
