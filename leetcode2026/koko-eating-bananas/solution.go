package koko_eating_bananas

import (
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, slices.Max(piles)
	for left < right {
		mid := left - (left-right)/2
		if getHours(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func getHours(piles []int, v int) int {
	hours := 0
	for _, pile := range piles {
		hours += pile / v
		if pile%v != 0 {
			hours++
		}
	}
	return hours
}
