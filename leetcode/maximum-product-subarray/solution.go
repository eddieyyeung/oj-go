package maximumproductsubarray

// maxProduct https://leetcode-cn.com/explore/interview/card/top-interview-questions-hard/60/dynamic-programming/154/
func maxProduct(nums []int) int {
	dp0 := nums[0] // 最大
	dp1 := nums[0] // 当前序列最大
	dp2 := nums[0] // 当前序列最小
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		k1 := max(dp1*n, dp2*n, n)
		k2 := min(dp1*n, dp2*n, n)
		k0 := max(dp0, k1)
		dp0, dp1, dp2 = k0, k1, k2
	}
	return dp0
}

func min(nums ...int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func max(nums ...int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
