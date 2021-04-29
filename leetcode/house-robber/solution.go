package house_robber

import "math"

func rob(nums []int) int {
	max := math.MinInt32
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		appendnum := 0
		for j := 0; j < i-1; j++ {
			if arr[j] > appendnum {
				appendnum = arr[j]
			}
		}
		arr[i] = nums[i] + appendnum
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
