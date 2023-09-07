package solution2

type UnionFind struct {
	Parent []int
	Rank   []int
}

func New(cap int) *UnionFind {
	parent := make([]int, cap)
	rank := make([]int, cap)
	for i := 0; i < cap; i++ {
		parent[i] = i
	}
	return &UnionFind{
		Parent: parent,
		Rank:   rank,
	}
}

func (uf UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		// path compression
		uf.Parent[x] = uf.Find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf UnionFind) Union(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}
	// union by rank
	if uf.Rank[x] > uf.Rank[y] {
		uf.Parent[y] = x
	} else if uf.Rank[x] < uf.Rank[y] {
		uf.Parent[x] = y
	} else {
		uf.Parent[x] = y
		uf.Rank[y]++
	}
}

func largestComponentSize(nums []int) int {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	uf := New(m + 1)
	for _, num := range nums {
		uf.Rank[num] = 1
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				uf.Union(i, num)
				uf.Union(num/i, num)
			}
		}
	}
	tm := make(map[int]int)
	var ans int
	for _, num := range nums {
		p := uf.Find(num)
		tm[p]++
		if tm[p] > ans {
			ans = tm[p]
		}
	}
	return ans
}
