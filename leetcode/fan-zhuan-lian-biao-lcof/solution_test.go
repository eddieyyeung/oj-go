package fanzhuanlianbiaolcof

import (
	"testing"
)

func Test_trainningPlan(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	trainningPlan(l)
}
