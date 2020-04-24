// Package solution1 ...
// https://leetcode-cn.com/problems/friend-circles/
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

func (ds *DisjointSet) CreateNode(i int) {
	if _, existed := ds.Father[i]; existed {
		return
	}
	ds.Father[i] = i
	ds.Rank[i] = 0
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

func findCircleNum(M [][]int) int {
	ds := NewDisjointSet()
	for i := 0; i < len(M); i++ {
		ds.CreateNode(i)
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 {
				ds.CreateNode(j)
				ds.Judge(i, j)
			}
		}
	}
	s := map[int]bool{}
	for k := range ds.Father {
		v := ds.GetFather(k)
		s[v] = true
	}
	return len(s)
}
