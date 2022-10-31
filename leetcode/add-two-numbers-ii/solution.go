package add_two_numbers_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a1 []int
	var a2 []int
	p1 := l1
	p2 := l2
	for p1 != nil {
		a1 = append(a1, p1.Val)
		p1 = p1.Next
	}
	for p2 != nil {
		a2 = append(a2, p2.Val)
		p2 = p2.Next
	}
	len1 := len(a1) - 1
	len2 := len(a2) - 1
	var result *ListNode
	carrybit := 0
	for len1 >= 0 && len2 >= 0 {
		sum := a1[len1] + a2[len2] + carrybit
		carrybit = sum / 10
		l := ListNode{
			Val:  sum % 10,
			Next: result,
		}
		result = &l
		len1--
		len2--
	}
	for len1 >= 0 {
		sum := a1[len1] + carrybit
		carrybit = sum / 10
		l := ListNode{
			Val:  sum % 10,
			Next: result,
		}
		result = &l
		len1--
	}
	for len2 >= 0 {
		sum := a2[len2] + carrybit
		carrybit = sum / 10
		l := ListNode{
			Val:  sum % 10,
			Next: result,
		}
		result = &l
		len2--
	}
	if carrybit != 0 {
		l := ListNode{
			Val:  carrybit,
			Next: result,
		}
		result = &l
	}
	return result
}
