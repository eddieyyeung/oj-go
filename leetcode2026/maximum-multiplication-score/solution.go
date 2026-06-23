package maximum_multiplication_score

import (
	"math"
)

func maxScore(a []int, b []int) int64 {
	cache := make([][]int64, len(a))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int64, len(b))
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = math.MinInt64
		}
	}

	var fn func(int, int) int64

	fn = func(i int, j int) int64 {
		if j < i {
			return -1e10 * 4
		}
		if i < 0 {
			return 0
		}
		if cache[i][j] != math.MinInt64 {
			return cache[i][j]
		}
		cache[i][j] = max(fn(i-1, j-1)+int64(a[i]*b[j]), fn(i, j-1))
		return cache[i][j]
	}

	ans := fn(len(a)-1, len(b)-1)
	return ans
}
