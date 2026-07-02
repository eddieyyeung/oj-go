package fruits_into_baskets_iii

import (
	"math/bits"
)

type seg []int

func (t seg) maintain(o int) {
	t[o] = max(t[o<<1], t[o<<1|1])
}

func (t seg) build(o, l, r int, a []int) {
	if l == r {
		t[o] = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m, a)
	t.build(o<<1|1, m+1, r, a)
	t.maintain(o)
}

func (t seg) findFirstAndUpdate(o, l, r, x int) int {
	if t[o] < x {
		return -1
	}
	if l == r {
		t[o] = -1
		return o
	}
	m := (l + r) >> 1
	i := t.findFirstAndUpdate(o<<1, l, m, x)
	if i < 0 {
		i = t.findFirstAndUpdate(o<<1|1, m+1, r, x)
	}
	t.maintain(o)
	return i
}

func newSegmentTree(a []int) seg {
	s := make(seg, 2<<bits.Len(uint(len(a))))
	s.build(1, 0, len(a)-1, a)
	return s
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	s := newSegmentTree(baskets)
	ans := 0
	for _, fruit := range fruits {
		if s.findFirstAndUpdate(1, 0, len(fruits)-1, fruit) == -1 {
			ans++
		}
	}
	return ans
}
