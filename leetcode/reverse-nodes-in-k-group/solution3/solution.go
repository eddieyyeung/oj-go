package solution3

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 1. 求出链表长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	// 2. 初始化哨兵节点
	sentry := &ListNode{Next: head}
	p0 := sentry
	var pre *ListNode
	cur = p0.Next
	// 3. 每 k 个长度的链表实现翻转，不足 k 长度直接退出
	for n >= k {
		n -= k
		pre = nil
		for i := 1; i <= k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 4. 处理翻转链表的首尾指针，并将 p0 指向下一个 p0 指针
		p0next := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = p0next
	}
	return sentry.Next
}
