// Package triangle ...
// https://leetcode-cn.com/problems/triangle/
package triangle

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = minF(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
		}
		dp[0] = dp[0] + triangle[i][0]
	}
	minT := math.MaxInt64
	for _, v := range dp {
		if v < minT {
			minT = v
		}
	}
	return minT
}

func minF(arr ...int) int {
	min := math.MaxInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
