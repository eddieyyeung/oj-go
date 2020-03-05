package maxproduct

import (
	"math"
)

// maxProduct https://leetcode-cn.com/explore/interview/card/top-interview-questions-hard/60/dynamic-programming/154/
func maxProduct(nums []int) int {
	max := 1
	min := 1
	finalMax := math.MinInt64
	for _, n := range nums {
		mmax := n * max
		mmin := n * min
		if mmax > mmin {
			max = mmax
			min = mmin
		} else {
			max = mmin
			min = mmax
		}
		if n > max {
			max = n
		}
		if min > n {
			min = n
		}
		if max > finalMax {
			finalMax = max
		}
	}
	return finalMax
}
