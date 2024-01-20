// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description
package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var q []*Node = []*Node{root}

	for len(q) != 0 {
		n := len(q)
		node := q[0]
		for i := 0; i < n; i++ {
			if i != 0 {
				node.Next = q[i]
				node = q[i]
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}
	return root
}
