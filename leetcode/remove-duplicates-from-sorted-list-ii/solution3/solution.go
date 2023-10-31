package solution3

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	sentry := &ListNode{Next: head}
	cur := sentry
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		// 1. 判断下一个节点和下下一个节点的值是否相等
		if cur.Next.Next.Val == val {
			// 2. 如果相等，则遍历删除相同值的节点
			for cur.Next != nil {
				if cur.Next.Val == val {
					cur.Next = cur.Next.Next
				} else {
					break
				}
			}
			// 3. 否则指向下一个指针
		} else {
			cur = cur.Next
		}
	}
	return sentry.Next
}
