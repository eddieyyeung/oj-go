package fruits_into_baskets_iii

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

type SegmentTree struct {
	Head *Node
}

func NewSegmentTree(a []int) *SegmentTree {
	st := &SegmentTree{Head: &Node{}}
	st.Build(st.Head, 0, len(a)-1, a)
	return st
}

func (st *SegmentTree) Maintain(node *Node) {
	lv, rv := -1, -1
	if node.Left != nil {
		lv = node.Left.Val
	}
	if node.Right != nil {
		rv = node.Right.Val
	}
	node.Val = max(lv, rv)
}

func (st *SegmentTree) Build(node *Node, l, r int, a []int) {
	if l == r {
		node.Val = a[l]
		return
	}
	mid := (l + r) >> 1

	node.Left = &Node{}
	st.Build(node.Left, l, mid, a)
	node.Right = &Node{}
	st.Build(node.Right, mid+1, r, a)
	st.Maintain(node)
}

func (st *SegmentTree) FindFirstAndUpdate(node *Node, x, l, r int) int {
	if node.Val < x {
		return -1
	}
	if l == r {
		v := node.Val
		node.Val = -1
		return v
	}
	mid := (l + r) >> 1
	rst := st.FindFirstAndUpdate(node.Left, x, l, mid)
	if rst < 0 {
		rst = st.FindFirstAndUpdate(node.Right, x, mid+1, r)
	}
	st.Maintain(node)
	return rst
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	st := NewSegmentTree(baskets)

	ans := 0
	for _, fruit := range fruits {
		if st.FindFirstAndUpdate(st.Head, fruit, 0, len(fruits)-1) == -1 {
			ans++
		}
	}
	return ans
}
