package solution

import (
	"math"
)

func minTransfers(transactions [][]int) int {
	n := 12
	cnt := make([]int, n)
	for _, transaction := range transactions {
		cnt[transaction[0]] -= transaction[2]
		cnt[transaction[1]] += transaction[2]
	}
	m := 1 << n
	dp := make([]int, m)
	for i := 1; i < m; i++ {
		var sum int
		for j := 0; j < n; j++ {
			if (i >> j & 1) == 1 {
				sum += cnt[j]
			}
		}
		if sum != 0 {
			dp[i] = math.MaxInt / 2
		} else {
			dp[i] = Count1(i) - 1
			for j := (i - 1) & i; j != 0; j = (j - 1) & i {
				dp[i] = MinInt(dp[i], dp[j]+dp[i^j])
			}
		}
	}
	return dp[m-1]
}

func MinInt(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// 计算出二进制数 x 中 1 的数量
func Count1(x int) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
