package maximumsubarray

func maxSubArray(nums []int) int {
	n := nums[0]
	max := n
	for i := 1; i < len(nums); i++ {
		m := nums[i] + n
		if m > nums[i] {
			n = m
		} else {
			n = nums[i]
		}
		if n > max {
			max = n
		}
	}
	return max
}
