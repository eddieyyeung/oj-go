// Package k_th_smallest_in_lexicographical_order ...
// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/
package k_th_smallest_in_lexicographical_order

import (
	"math"
)

func findKthNumber(n int, k int) int {
	cur := 1
	k--
	for k > 0 {
		step := 0
		first := cur
		last := cur + 1
		for first <= n {
			step += min(n+1, last) - first
			first *= 10
			last *= 10
		}
		if k < step {
			k--
			cur *= 10
		} else {
			k -= step
			cur++
		}
	}
	return cur
}

func min(a ...int) int {
	min := math.MaxInt64
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
