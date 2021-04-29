package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	p := head
	q := newHead
	for p.Next != nil {
		if p.Next == head {
			break
		}
		newNode := &Node{
			Val:    p.Next.Val,
			Next:   nil,
			Random: nil,
		}
		p = p.Next
		q.Next = newNode
		q = q.Next
	}
	p = head
	q = newHead
	p1 := head
	q1 := newHead
	for p != nil {
		if p.Random != nil {
			p1 = head
			q1 = newHead
			for p1 != nil {
				if p.Random == p1 {
					q.Random = q1
				}
				p1 = p1.Next
				q1 = q1.Next
				if p1 == head {
					break
				}
			}
		}
		p = p.Next
		q = q.Next
		if p == head {
			break
		}
	}
	return newHead
}
