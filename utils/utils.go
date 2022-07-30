package utils

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateNode(arr []*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   *arr[0],
		Left:  nil,
		Right: nil,
	}
	l := list.New()
	l.PushBack(root)
	for i := 1; i < len(arr); i++ {
		e := l.Front()
		p := e.Value.(*TreeNode)
		var left, right *TreeNode
		if arr[i] != nil {
			left = &TreeNode{
				Val:   *arr[i],
				Left:  nil,
				Right: nil,
			}
			l.PushBack(left)
		}
		if i+1 < len(arr) {
			i = i + 1
			if arr[i] != nil {
				right = &TreeNode{
					Val:   *arr[i],
					Left:  nil,
					Right: nil,
				}
				l.PushBack(right)
			}
		}
		p.Left = left
		p.Right = right
		l.Remove(e)
	}
	return root
}

func MinInt(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func MaxInt(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Count1(x int) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
