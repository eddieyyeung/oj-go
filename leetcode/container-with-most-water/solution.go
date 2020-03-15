// Packge container_with_most_water ...
// https://leetcode-cn.com/problems/container-with-most-water/
package container_with_most_water

import "math"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		if n := minf(height[l], height[r]) * (r - l); n > max {
			max = n
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func minf(a ...int) int {
	min := math.MaxInt64
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
