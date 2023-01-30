package merge_in_between_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	i := 0
	p := list1
	var lprev, r, prev *ListNode
	for p != nil {
		if i == a {
			lprev = prev
		}
		if i == b {
			r = p
		}
		prev = p
		p = p.Next
		i++
	}
	var tail *ListNode
	p = list2
	for {
		if p.Next == nil {
			tail = p
			break
		}
		p = p.Next
	}
	if lprev == nil {
		tail.Next = r.Next
		r.Next = nil
		return list2
	}
	lprev.Next = list2
	tail.Next = r.Next
	r.Next = nil
	return list1
}
