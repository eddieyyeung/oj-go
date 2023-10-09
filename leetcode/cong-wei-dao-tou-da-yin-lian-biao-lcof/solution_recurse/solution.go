package solutionrecurse

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBookList(head *ListNode) []int {
	var ans []int
	recur(head, &ans)
	return ans
}

func recur(p *ListNode, ans *[]int) {
	if p != nil {
		recur(p.Next, ans)
		*ans = append(*ans, p.Val)
	}
}
