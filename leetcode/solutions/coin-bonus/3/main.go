package solution

import (
	"container/list"
)

type Node struct {
	Left  *Node
	Right *Node
	Value uint64
	Idx   int
}

type SegmentTree struct {
	Root    *Node
	Members *[]Member
}

type Member struct {
	SubFollowers *list.List
	Index        int
	Start        int
	End          int
	Idx          int
	SubCoins     uint64
}

func (st *SegmentTree) BuildLeft(node *Node) *Node {
	if node.Left == nil {
		node.Left = &Node{nil, nil, 0, 0}
	}
	return node.Left
}

func (st *SegmentTree) BuildRight(node *Node) *Node {
	if node.Right == nil {
		node.Right = &Node{nil, nil, 0, 0}
	}
	return node.Right
}

func (st *SegmentTree) Update(node *Node, start int, end int, l int, r int, val uint64) {
	if start == end {
		// Leaf node
		node.Value += val
	} else {
		mid := (start + end) / 2
		if l <= mid {
			st.Update(st.BuildLeft(node), start, mid, l, r, val)
		}
		if mid+1 <= r {
			st.Update(st.BuildRight(node), mid+1, end, l, r, val)
		}
		node.Value = st.BuildLeft(node).Value + st.BuildRight(node).Value
	}
}

func (st *SegmentTree) Query(node *Node, start int, end int, l int, r int) uint64 {
	if node == nil {
		return 0
	}
	if l <= start && end <= r {
		return node.Value
	}
	mid := (start + end) / 2
	p1 := st.Query(node.Left, start, mid, l, r)
	p2 := st.Query(node.Right, mid+1, end, l, r)
	return p1 + p2
}

func (st *SegmentTree) Print() {
	l := list.New()
	l.PushBack(st.Root)
	for l.Len() != 0 {
		e := l.Front()
		node, _ := e.Value.(*Node)
		l.Remove(e)
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
}

func preorderTraversal(arr []*Member, p *Member, node *int) {
	arr[*node] = p
	p.Start = *node
	p.Idx = *node
	*node++
	for e := p.SubFollowers.Front(); e != nil; e = e.Next() {
		m, _ := e.Value.(*Member)
		preorderTraversal(arr, m, node)
	}
	p.End = *node - 1
}

var ms []Member
var arr []*Member

func bonus(n int, leadership [][]int, operations [][]int) []int {
	ms = make([]Member, n)
	result := []int{}
	for i := 0; i < n; i++ {
		ms[i] = Member{list.New(), i, 0, 0, 0, 0}
	}
	for i := 0; i < len(leadership); i++ {
		ms[leadership[i][0]-1].SubFollowers.PushBack(&ms[leadership[i][1]-1])
	}
	root := &ms[0]
	arr = make([]*Member, n)
	idx := 0
	preorderTraversal(arr, root, &idx)
	st := &SegmentTree{&Node{nil, nil, 0, 0}, &ms}
	l := list.New()
	for i := 0; i < len(operations); i++ {
		o1 := operations[i][0]
		switch o1 {
		case 1:
			o2 := operations[i][1]
			m := ms[o2-1]
			coins := uint64(operations[i][2])
			st.Update(st.Root, 0, n-1, m.Idx, m.Idx, coins)
		case 2:
			o2 := operations[i][1]
			m := ms[o2-1]
			coins := uint64(operations[i][2])
			m.SubCoins += coins
			l.PushBack(&m)

			// st.Update(st.Root, 0, n-1, m.Idx, m.Idx, coins)
		case 3:
			o2 := operations[i][1]
			m := ms[o2-1]
			earr := []*list.Element{}
			for e := l.Front(); e != nil; e = e.Next() {
				t, _ := e.Value.(*Member)
				if t.SubCoins != 0 {
					if !(t.Start > m.End || t.End < m.Start) {
						st.Update(st.Root, 0, n-1, t.Start, t.End, t.SubCoins)
						t.SubCoins = 0
						earr = append(earr, e)
					}
				}
			}
			for _, v := range earr {
				l.Remove(v)
			}

			result = append(result, int(st.Query(st.Root, 0, n-1, m.Start, m.End)%1000000007))
		}
	}
	return result
}

func Run(n int, leadership [][]int, operations [][]int) []int {
	return bonus(n, leadership, operations)
}
