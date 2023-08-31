package gupiaodezuidalirunlcof

import "math"

func maxProfit(prices []int) int {
	m := 0
	s := math.MaxInt
	for _, price := range prices {
		if profit := price - s; profit > m {
			m = profit
		}
		if price < s {
			s = price
		}
	}
	return m
}
