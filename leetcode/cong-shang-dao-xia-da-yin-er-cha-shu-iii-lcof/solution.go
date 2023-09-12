package congshangdaoxiadayinerchashuiiilcof

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	l := list.New()
	l.PushBack(root)
	var ans [][]int
	i := 1
	for l.Len() > 0 {
		lb := list.New()
		var level []int
		if i == 1 {
			for e := l.Front(); e != nil; e = e.Next() {
				node := e.Value.(*TreeNode)
				if node != nil {
					lb.PushBack(node.Left)
					lb.PushBack(node.Right)
					level = append(level, node.Val)
				}
			}
		} else {
			for e := l.Back(); e != nil; e = e.Prev() {
				node := e.Value.(*TreeNode)
				if node != nil {
					lb.PushFront(node.Right)
					lb.PushFront(node.Left)
					level = append(level, node.Val)
				}
			}
		}
		i = i ^ 1
		if len(level) != 0 {
			ans = append(ans, level)
		}
		l = lb
	}
	return ans
}
