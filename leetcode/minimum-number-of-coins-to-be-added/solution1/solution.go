// 2952. 需要添加的硬币的最小数量
// https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/description/
package solution

import (
	"sort"
)

func minimumAddedCoins(coins []int, target int) int {
	n := len(coins)
	sort.Ints(coins)
	var ans int
	var sum int
	i := 0
	for sum < target {
		if i < n && coins[i] <= sum+1 {
			sum += coins[i]
			i++
		} else {
			sum += sum + 1
			ans++
		}
	}
	return ans
}
