// 2487. 从链表中移除节点
// https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	var q []*ListNode
	p := head
	for p != nil {
		for len(q) > 0 && q[len(q)-1].Val < p.Val {
			q = q[:len(q)-1]
		}
		q = append(q, p)
		p = p.Next
	}
	sentry := &ListNode{}
	p = sentry
	for _, node := range q {
		node.Next = nil
		p.Next = node
		p = p.Next
	}
	return sentry.Next
}
