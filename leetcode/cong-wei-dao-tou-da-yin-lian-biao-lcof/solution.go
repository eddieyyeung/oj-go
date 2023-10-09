package congweidaotoudayinlianbiaolcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBookList(head *ListNode) []int {
	n := 10000
	ans := make([]int, n)
	p := head
	for p != nil {
		n--
		ans[n] = p.Val
		p = p.Next
	}
	return ans[n:]
}
