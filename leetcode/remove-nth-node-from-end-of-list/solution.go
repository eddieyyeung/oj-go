package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{
		Next: head,
	}
	// 1. 初始化 p0, p1 指向哨兵节点
	p0, p1 := sentry, sentry
	// 2. p1 先走 n 步
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	// 3. p0, p1 同时继续走，持续到 p1.Next == nil 为止，
	//    则 p0, p1 刚好相差 n 步，p0 的下一个节点则为待删除节点
	for p1.Next != nil {
		p1 = p1.Next
		p0 = p0.Next
	}
	// 4. 删除
	nxt := p0.Next
	p0.Next = p0.Next.Next
	nxt.Next = nil
	return sentry.Next
}
