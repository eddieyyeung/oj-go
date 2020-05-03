// Package edit_distance ...
// https://leetcode-cn.com/problems/edit-distance/
package edit_distance

import "math"

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word1)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len(word2)][len(word1)]
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
