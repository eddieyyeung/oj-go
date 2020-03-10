// Package howmanynumbersaresmallerthanthecurrentnumber ...
// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	r := []int{}
	for i := 0; i < len(nums); i++ {
		smaller := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[j] < nums[i] {
				smaller++
			}
		}
		r = append(r, smaller)
	}
	return r
}
