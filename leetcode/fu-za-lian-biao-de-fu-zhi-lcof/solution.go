package fuzalianbiaodefuzhilcof

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var copyHead *Node = &Node{}
	m := make(map[*Node]*Node)
	p := head
	q := copyHead
	for p != nil {
		cpNode := &Node{
			Val: p.Val,
		}
		q.Next = cpNode
		q = q.Next
		m[p] = cpNode
		p = p.Next
	}
	p = head
	q = copyHead.Next
	for p != nil {
		if p.Random == nil {
			p = p.Next
			q = q.Next
			continue
		}
		q.Random = m[p.Random]
		q = q.Next
		p = p.Next
	}
	return copyHead.Next
}
