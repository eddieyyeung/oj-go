package segmenttreelink

type NumArray struct {
	Head *SegmentTreeLinkNode
	Size int
}

type SegmentTreeLinkNode struct {
	Value int
	Left  *SegmentTreeLinkNode
	Right *SegmentTreeLinkNode
}

func Constructor(nums []int) NumArray {
	stl := NumArray{
		Head: &SegmentTreeLinkNode{},
		Size: len(nums),
	}
	stl.build(stl.Head, 0, stl.Size-1, nums)
	return stl
}

func (stl *NumArray) build(node *SegmentTreeLinkNode, s, e int, nums []int) {
	if s == e {
		node.Value = nums[s]
		return
	}
	if node.Left == nil {
		node.Left = &SegmentTreeLinkNode{}
	}
	if node.Right == nil {
		node.Right = &SegmentTreeLinkNode{}
	}
	m := s + ((e - s) >> 1)
	stl.build(node.Left, s, m, nums)
	stl.build(node.Right, m+1, e, nums)
	node.Value = node.Left.Value + node.Right.Value
}

func (stl *NumArray) Update(index int, value int) {
	stl.update(stl.Head, 0, stl.Size-1, index, value)
}

func (stl *NumArray) update(node *SegmentTreeLinkNode, s, e, index, value int) {
	if s == e {
		node.Value = value
		return
	}
	m := s + ((e - s) >> 1)
	if index <= m {
		stl.update(node.Left, s, m, index, value)
	} else {
		stl.update(node.Right, m+1, e, index, value)
	}
	node.Value = node.Left.Value + node.Right.Value
}

func (stl *NumArray) SumRange(left int, right int) int {
	return stl.sumrange(stl.Head, 0, stl.Size-1, left, right)
}

func (stl *NumArray) sumrange(node *SegmentTreeLinkNode, s, e, left, right int) int {
	if s == left && e == right {
		return node.Value
	}
	m := s + ((e - s) >> 1)
	if right <= m {
		return stl.sumrange(node.Left, s, m, left, right)
	}
	if left > m {
		return stl.sumrange(node.Right, m+1, e, left, right)
	}
	return stl.sumrange(node.Left, s, m, left, m) + stl.sumrange(node.Right, m+1, e, m+1, right)
}
