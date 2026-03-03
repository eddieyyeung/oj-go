package maximum_total_damage_with_spell_casting

import (
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	m := map[int]int64{}
	for _, x := range power {
		m[x] += int64(x)
	}
	var nums []int
	for num := range m {
		nums = append(nums, num)
	}
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int64, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	var dfs func(i int) int64
	dfs = func(i int) int64 {
		if i < 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		j := i - 1
		for j >= 0 && nums[j] >= nums[i]-2 {
			j--
		}
		dp[i] = max(dfs(i-1), m[nums[i]]+dfs(j))

		return dp[i]
	}
	return dfs(n - 1)
}
