package ways_to_make_a_fair_array

func waysToMakeFair(nums []int) int {
	prevEven := 0
	prevOdd := 0
	suffEven := 0
	suffOdd := 0
	for i := range nums {
		if i == 0 {
			continue
		}
		if i%2 == 0 {
			suffOdd += nums[i]
		} else {
			suffEven += nums[i]
		}
	}
	result := 0
	if prevEven+suffEven == prevOdd+suffOdd {
		result++
	}
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			prevEven += nums[i]
			suffEven -= nums[i+1]
		} else {
			prevOdd += nums[i]
			suffOdd -= nums[i+1]
		}
		if prevEven+suffEven == prevOdd+suffOdd {
			result++
		}
	}
	return result
}
