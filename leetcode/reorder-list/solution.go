package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 1. 利用快慢指针求出中间节点
	mid := middleNode(head)
	// 2. 反转 mid 之后的链表为 head2，注意 mid.Next = nil
	head2 := reserveList(mid.Next)
	mid.Next = nil
	p1, p2 := head, head2
	// 3. 遍历 head2 链表，在 head 链表插入节点
	for p2 != nil {
		nxt := p1.Next
		p1.Next = p2
		p2 = p2.Next
		p1.Next.Next = nxt
		p1 = nxt
	}
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reserveList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
