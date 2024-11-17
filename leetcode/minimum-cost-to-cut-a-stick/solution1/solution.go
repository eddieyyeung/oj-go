// https://leetcode.cn/problems/minimum-cost-to-cut-a-stick
package minimumcosttocutastick

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i] < cuts[j]
	})
	states := make([][2]int, len(cuts))
	for i := 0; i < len(cuts); i++ {
		states[i][0] = cuts[i]
	}
	type item struct {
		l int
		r int
	}

	dp := make(map[item]int)
	var ans int = math.MaxInt

	var fn func(l, r int) int

	fn = func(l, r int) int {
		if _, ok := dp[item{l, r}]; ok {
			return dp[item{l, r}]
		}

		var cost int = math.MaxInt

		for i := 0; i < len(states); i++ {
			if states[i][1] == 0 && states[i][0] > l && states[i][0] < r {
				states[i][1] = 1
				cost = min(cost, r-l+fn(l, states[i][0])+fn(states[i][0], r))
				states[i][1] = 0
			}
		}
		if cost == math.MaxInt {
			cost = 0
		}
		dp[item{l, r}] = cost
		return cost
	}
	ans = fn(0, n)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
