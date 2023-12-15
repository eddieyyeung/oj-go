// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description/
package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var p *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow_next := slow.Next
		slow.Next = p
		p = slow
		slow = slow_next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != p.Val {
			return false
		}
		slow = slow.Next
		p = p.Next
	}
	return true
}
