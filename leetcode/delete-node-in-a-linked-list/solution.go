package deletenodeinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 237. 删除链表中的节点
// https://leetcode.cn/problems/delete-node-in-a-linked-list/description/
// 无法通过上一个节点指向下一个节点的方式删除节点
// 解决思路：复制下一个节点的值，并修改 Next 指针
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
