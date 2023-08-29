package removeduplicatesfromsortedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodePosition struct {
	Prev    *ListNode
	Ptr     *ListNode
	Deleted bool
}

// The list is a sorted linked list, but this solution doesn't consider it, so uses map to deduplicate nodes.
// There is another optimal solution to solve this problem as a sorted linked list.
func deleteDuplicates(head *ListNode) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	m := make(map[int]NodePosition)

	p := sentry
	for p.Next != nil {
		prev := p
		p = p.Next
		if pos, ok := m[p.Val]; ok {
			pprev := prev
			if !pos.Deleted {
				if prev == pos.Ptr {
					pprev = pos.Prev
				}
				pos.Prev.Next = pos.Ptr.Next
				pos.Ptr.Next = nil
				pos.Deleted = true
			}
			pprev.Next = p.Next
			p.Next = nil
			p = pprev

		} else {
			m[p.Val] = NodePosition{
				Prev:    prev,
				Ptr:     p,
				Deleted: false,
			}
		}
	}
	return sentry.Next
}
