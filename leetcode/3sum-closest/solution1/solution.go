package solution

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	var ans int = int(1e9)
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			sub := sum - target
			if sub == 0 {
				return sum
			}
			if sub > 0 {
				r--
			} else {
				sub = -sub
				l++
			}
			if sub < abs(ans, target) {
				ans = sum
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return a - b
}
