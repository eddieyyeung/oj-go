// Package solution1 ...
// https://leetcode-cn.com/problems/longest-consecutive-sequence/
package solution1

type DisjointSet struct {
	Father map[int]int
	Rank   map[int]int
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		Father: map[int]int{},
		Rank:   map[int]int{},
	}
}

func (ds *DisjointSet) GetFather(i int) int {
	if ds.Father[i] == i {
		return i
	}
	ds.Father[i] = ds.GetFather(ds.Father[i])
	return ds.Father[i]
}

func (ds *DisjointSet) Judge(x int, y int) {
	fx := ds.GetFather(x)
	fy := ds.GetFather(y)
	if ds.Rank[fx] > ds.Rank[fy] {
		ds.Father[fy] = fx
	} else {
		ds.Father[fx] = fy
		if ds.Rank[fx] == ds.Rank[fy] {
			ds.Rank[fy]++
		}
	}
}

func longestConsecutive(nums []int) int {
	countM := map[int]int{}
	ds := NewDisjointSet()
	l := 0
	for _, num := range nums {
		if _, existed := ds.Father[num]; !existed {
			ds.Father[num] = num
			countM[num] = 1
		} else {
			continue
		}
		fn := ds.GetFather(num)
		s := countM[fn]
		if n1, existed := ds.Father[num-1]; existed {
			fn1 := ds.GetFather(n1)
			sn1 := countM[fn1]
			s += sn1
			ds.Judge(fn, fn1)
			countM[ds.GetFather(fn)] = s
		}
		if n1, existed := ds.Father[num+1]; existed {
			fn1 := ds.GetFather(n1)
			sn1 := countM[fn1]
			s += sn1
			ds.Judge(fn, fn1)
			countM[ds.GetFather(fn)] = s
		}
		if s > l {
			l = s
		}
	}
	return l
}
