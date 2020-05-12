// Package largest_component_size_by_common_factor ...
// https://leetcode-cn.com/problems/largest-component-size-by-common-factor/
package largest_component_size_by_common_factor

type DisjointSet struct {
	Father []int
	Rank   []int
}

func NewDisjointSet(cap int) *DisjointSet {
	father := make([]int, cap)
	rank := make([]int, cap)
	for i := 0; i < cap; i++ {
		father[i] = i
	}
	return &DisjointSet{
		Father: father,
		Rank:   rank,
	}
}

func (ds *DisjointSet) GetFather(i int) int {
	if ds.Father[i] != i {
		ds.Father[i] = ds.GetFather(ds.Father[i])
	}
	return ds.Father[i]
}

func (ds *DisjointSet) Union(x, y int) {
	fx := ds.GetFather(x)
	fy := ds.GetFather(y)
	if fx == fy {
		return
	}
	if ds.Rank[fx] > ds.Rank[fy] {
		ds.Father[fy] = fx
	} else if ds.Rank[fx] < ds.Rank[fy] {
		ds.Father[fx] = fy
	} else {
		ds.Father[fx] = fy
		ds.Rank[fy]++
	}
}

func (ds *DisjointSet) IsSame(i int, j int) bool {
	return ds.GetFather(i) == ds.GetFather(j)
}

func findPrimes(x int) []int {
	m := map[int]bool{}
	var a []int
	p := x
	for i := 2; i*i <= p; i++ {
		for p%i == 0 {
			if _, ok := m[i]; !ok {
				m[i] = true
				a = append(a, i)
			}
			p = p / i
		}
	}
	if p > 1 {
		a = append(a, p)
	}
	return a
}

func largestComponentSize(A []int) int {
	pm := map[int]int{}
	var pa []int
	var pas [][]int
	for i := 0; i < len(A); i++ {
		ps := findPrimes(A[i])
		for _, p := range ps {
			if _, ok := pm[p]; !ok {
				pa = append(pa, p)
				pm[p] = len(pa) - 1
			}
		}
		pas = append(pas, ps)
	}

	ds := NewDisjointSet(len(pa))
	m := map[int]int{}
	for i := 0; i < len(pas); i++ {
		item := pas[i]
		if len(item) == 0 {
			continue
		}
		first := item[0]
		for j := 1; j < len(item); j++ {
			ds.Union(pm[first], pm[item[j]])
		}
	}
	c := 0
	for i := 0; i < len(pas); i++ {
		item := pas[i]
		if len(item) == 0 {
			continue
		}
		fi := ds.GetFather(pm[item[0]])
		m[fi]++
		if m[fi] > c {
			c = m[fi]
		}
	}
	return c
}
