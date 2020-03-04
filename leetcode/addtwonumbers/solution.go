package addtwonumbers

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	root := &ListNode{
		sum % 10,
		nil,
	}
	lnp := root
	inc := sum / 10

	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + inc
		inc = sum / 10
		lnp.Next = &ListNode{
			sum % 10,
			nil,
		}
		lnp = lnp.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + inc
		inc = sum / 10
		lnp.Next = &ListNode{
			sum % 10,
			nil,
		}
		lnp = lnp.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + inc
		inc = sum / 10
		lnp.Next = &ListNode{
			sum % 10,
			nil,
		}
		lnp = lnp.Next
		l2 = l2.Next
	}

	if inc != 0 {
		lnp.Next = &ListNode{
			inc,
			nil,
		}
	}
	return root
}
