// Package sqrtx ...
// https://leetcode-cn.com/problems/sqrtx/
package sqrtx

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left := 1
	right := x / 2
	for left <= right {
		mid := (left + right) / 2
		mulMid := mid * mid
		if mulMid == x {
			return mid
		}
		if x < mulMid {
			if x > (mid-1)*(mid-1) {
				return mid - 1
			}
			right = mid - 1
		}
		if x > mulMid {
			if x < (mid+1)*(mid+1) {
				return mid
			}
			left = mid + 1
		}
	}
	return 1
}
