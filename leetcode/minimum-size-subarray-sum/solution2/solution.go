package solution2

// 双指针
func minSubArrayLen(target int, nums []int) int {
	var ans int = len(nums) + 1
	var sum int
	var j int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum-nums[j] >= target {
			sum -= nums[j]
			j++
		}
		if sum >= target && i-j+1 < ans {
			ans = i - j + 1
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
