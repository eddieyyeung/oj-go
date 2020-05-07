// Package solution1 ...
// https://leetcode-cn.com/problems/russian-doll-envelopes/
package solution1

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		e1 := envelopes[i]
		e2 := envelopes[j]
		if e1[0] == e2[0] {
			return e1[1] < e2[1]
		} else {
			return e1[0] < e2[0]
		}
	})
	dp := make([]int, len(envelopes))
	dp[0] = 1
	maxLength := 1
	for i := 1; i < len(envelopes); i++ {
		ei := envelopes[i]
		l := 1
		for j := 0; j < i; j++ {
			ej := envelopes[j]
			d := dp[j]
			if ei[0] > ej[0] && ei[1] > ej[1] {
				l = maxF(d+1, l)
			}
		}
		dp[i] = l
		if l > maxLength {
			maxLength = l
		}
	}
	return maxLength
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
