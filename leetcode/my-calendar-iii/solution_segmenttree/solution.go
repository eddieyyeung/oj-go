package mycalendariii

type SegmentTreeNode struct {
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
	Value int // 当前节点的值
	Add   int // 懒惰标记
}

type SegmentTree struct {
	Head *SegmentTreeNode
	Size int
}

func NewSegmentTree(size int) *SegmentTree {
	return &SegmentTree{
		Head: &SegmentTreeNode{},
		Size: size,
	}
}

// leftNum 和 rightNum 表示左右孩子区间的叶子节点数量
// 因为如果是「加减」更新操作的话，需要用懒惰标记的值✖️叶子节点的数量
func (st *SegmentTree) pushDown(node *SegmentTreeNode, leftNum int, rightNum int) {
	// 动态开点
	if node.Left == nil {
		node.Left = &SegmentTreeNode{}
	}
	if node.Right == nil {
		node.Right = &SegmentTreeNode{}
	}
	// 如果 add 为 0，表示没有标记
	if node.Add == 0 {
		return
	}
	// 注意：当前节点加上标记值✖️该子树所有叶子节点的数量
	node.Left.Value += node.Add
	node.Right.Value += node.Add
	// 把标记下推给孩子节点
	// 对区间进行「加减」的更新操作，下推懒惰标记时需要累加起来，不能直接覆盖
	node.Left.Add += node.Add
	node.Right.Add += node.Add
	// 取消当前节点标记
	node.Add = 0
}

// 在区间 [start, end] 中更新区间 [l, r] 的值，将区间 [l, r] + val
// 对于上面的例子，应该这样调用该函数：Update(root, 0, 4, 2, 4, 1)
func (st *SegmentTree) Update(node *SegmentTreeNode, start int, end int, left int, right int, val int) {
	// 找到满足要求的区间
	if left <= start && end <= right {
		// 区间节点加上更新值
		// 注意：需要 * 该子树所有叶子节点
		node.Value += val
		node.Add += val
		return
	}
	mid := (start + end) >> 1
	// 下推标记
	// mid - start + 1：表示左孩子区间叶子节点数量
	// end - mid：表示右孩子区间叶子节点数量
	st.pushDown(node, mid-start+1, end-mid)
	// [start, mid] 和 [l, r] 可能有交集，遍历左孩子区间
	if left <= mid {
		st.Update(node.Left, start, mid, left, right, val)
	}
	// [mid + 1, end] 和 [l, r] 可能有交集，遍历右孩子区间
	if right > mid {
		st.Update(node.Right, mid+1, end, left, right, val)
	}
	// 向上更新
	st.pushUp(node)
}

func (st *SegmentTree) pushUp(node *SegmentTreeNode) {
	node.Value = max(node.Left.Value, node.Right.Value)
}

func (st *SegmentTree) Query(node *SegmentTreeNode, start int, end int, left int, right int) int {
	// 区间 [l ,r] 完全包含区间 [start, end]
	// 例如：[2, 4] = [2, 2] + [3, 4]，当 [start, end] = [2, 2] 或者 [start, end] = [3, 4]，直接返回
	// fmt.Println(start, end, left, right, node.Value, node.Add)
	if left <= start && end <= right {
		return node.Value
	}
	// 把当前区间 [start, end] 均分得到左右孩子的区间范围
	// node 左孩子区间 [start, mid]
	// node 左孩子区间 [mid + 1, end]
	mid := (start + end) >> 1
	ans := 0
	// 下推标记
	st.pushDown(node, mid-start+1, end-mid)
	// [start, mid] 和 [l, r] 可能有交集，遍历左孩子区间
	if left <= mid {
		ans = st.Query(node.Left, start, mid, left, right)
	}
	// [mid + 1, end] 和 [l, r] 可能有交集，遍历右孩子区间
	if right > mid {
		ans = max(ans, st.Query(node.Right, mid+1, end, left, right))
	}
	// ans 把左右子树的结果都累加起来了，与树的后续遍历同理
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MyCalendarThree struct {
	ST *SegmentTree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		ST: NewSegmentTree(1e9),
	}
}

func (mc *MyCalendarThree) Book(startTime int, endTime int) int {
	mc.ST.Update(mc.ST.Head, 0, mc.ST.Size-1, startTime, endTime-1, 1)
	return mc.ST.Query(mc.ST.Head, 0, mc.ST.Size-1, 0, mc.ST.Size-1)
}
