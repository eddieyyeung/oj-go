// https://leetcode.cn/problems/minimum-cost-to-cut-a-stick
package minimumcosttocutastick

import (
	"sort"
)

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i] < cuts[j]
	})
	m := len(cuts)
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := i + 2; j < len(dp[i]); j++ {
			dp[i][j] = 1e9
		}
	}
	for i := 2; i < m; i++ {
		for j := 0; j < m-i; j++ {
			for k := j + 1; k < j+i; k++ {
				dp[j][j+i] = min(dp[j][j+i], dp[j][k]+dp[k][j+i]+cuts[j+i]-cuts[j])
			}
		}
	}
	return dp[0][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
