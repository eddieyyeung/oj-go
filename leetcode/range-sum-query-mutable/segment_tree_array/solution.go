package segmenttreearray

type NumArray []int

func Constructor(nums []int) NumArray {
	str := make(NumArray, len(nums)*4)
	str.build(0, 0, len(nums)-1, nums)
	return str
}

func (str NumArray) Len() int {
	return len(str) / 4
}

func (str NumArray) build(node, s, e int, nums []int) {
	if s == e {
		str[node] = nums[s]
		return
	}
	m := s + ((e - s) >> 1)
	str.build(node*2+1, s, m, nums)
	str.build(node*2+2, m+1, e, nums)
	str[node] = str[node*2+1] + str[node*2+2]
}

func (str NumArray) update(node, s, e, index, value int) {
	if s == e {
		str[node] = value
		return
	}
	m := s + ((e - s) >> 1)
	if index <= m {
		str.update(node*2+1, s, m, index, value)
	} else {
		str.update(node*2+2, m+1, e, index, value)
	}
	str[node] = str[node*2+1] + str[node*2+2]
}

func (str NumArray) Update(index int, val int) {
	str.update(0, 0, str.Len()-1, index, val)
}

func (str NumArray) sumrange(node, s, e, left, right int) int {
	if s == left && e == right {
		return str[node]
	}
	m := s + ((e - s) >> 1)
	if right <= m {
		return str.sumrange(node*2+1, s, m, left, right)
	}
	if left > m {
		return str.sumrange(node*2+2, m+1, e, left, right)
	}
	return str.sumrange(node*2+1, s, m, left, m) + str.sumrange(node*2+2, m+1, e, m+1, right)
}

func (str NumArray) SumRange(left int, right int) int {
	return str.sumrange(0, 0, str.Len()-1, left, right)
}
