// Package candy ...
// https://leetcode-cn.com/problems/candy/
package candy

import "math"

func candy(ratings []int) int {
	arr := make([]int, len(ratings))
	arr[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 1
		}
	}
	sum := arr[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			arr[i] = maxF(arr[i], arr[i+1]+1)
		} else {
			arr[i] = maxF(arr[i], 1)
		}
		sum += arr[i]
	}
	return sum
}

func maxF(arr ...int) int {
	max := math.MinInt64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
