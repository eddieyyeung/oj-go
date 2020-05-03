// Package super_egg_drop ...
// https://leetcode-cn.com/problems/super-egg-drop/
package super_egg_drop

import (
	"math"
)

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i < N+1; i++ {
		dp[1][i] = i
	}
	for i := 2; i <= K; i++ {
		dp[i][1] = 1
		for j := 2; j <= N; j++ {
			m := dp[i-1][j]
			left := 1
			right := j
			for left <= right {
				mid := (right + left) / 2
				if dp[i-1][mid-1] > dp[i][j-mid] {
					right = mid - 1
					m = min(dp[i-1][mid-1]+1, m)
				} else {
					left = mid + 1
					m = min(dp[i][j-mid]+1, m)
				}
			}
			dp[i][j] = m
		}
	}
	return dp[K][N]
}

func min(arr ...int) int {
	min := math.MaxInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
