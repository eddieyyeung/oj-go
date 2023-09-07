package solution2

type UnionFind struct {
	Parent []int
	Rank   []int
}

func newUnionFind(n int) UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return UnionFind{
		Parent: parent,
		Rank:   rank,
	}
}

func (uf UnionFind) find(x int) int {
	if uf.Parent[x] != x {
		// Path Compression
		uf.Parent[x] = uf.find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf UnionFind) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	// Union by rank
	if uf.Rank[x] > uf.Rank[y] {
		uf.Parent[y] = x
	} else if uf.Rank[y] > uf.Rank[x] {
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
	uf := newUnionFind(m + 1)
	for _, num := range nums {
		uf.Rank[num] = 1
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				uf.merge(i, num)
				uf.merge(num/i, num)
			}
		}
	}
	tm := make(map[int]int)
	var ans int
	for _, num := range nums {
		p := uf.find(num)
		tm[p]++
		if tm[p] > ans {
			ans = tm[p]
		}
	}
	return ans
}
