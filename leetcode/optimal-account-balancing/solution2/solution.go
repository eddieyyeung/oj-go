package solution2

import (
	"math"
)

var result = math.MaxInt

func minTransfers(transactions [][]int) int {
	n := 12
	result = math.MaxInt
	cnt := make([]int, n)
	for _, transaction := range transactions {
		cnt[transaction[0]] -= transaction[2]
		cnt[transaction[1]] += transaction[2]
	}
	var aarr []int
	var barr []int
	for _, c := range cnt {
		if c > 0 {
			aarr = append(aarr, c)
		} else if c < 0 {
			barr = append(barr, -c)
		}
	}
	//sort.Ints(aarr)
	//sort.Ints(barr)
	dfs(aarr, barr, 0)
	return result
}

func dfs(aarr, barr []int, ans int) {
	if ans >= result {
		return
	}
	isAllZero := true
	for _, v := range aarr {
		if v != 0 {
			isAllZero = false
			break
		}
	}
	if isAllZero {
		result = ans
		return
	}
	mn := math.MaxInt
	var belong int
	var belong_idx int
	for i := 0; i < len(aarr); i++ {
		if aarr[i] > 0 && aarr[i] < mn {
			mn = aarr[i]
			belong = 0
			belong_idx = i
		}
	}
	for i := 0; i < len(barr); i++ {
		if barr[i] > 0 && barr[i] < mn {
			mn = barr[i]
			belong = 1
			belong_idx = i
		}
	}
	if belong == 0 {
		aarr[belong_idx] = 0
		for i := 0; i < len(barr); i++ {
			if barr[i] > 0 {
				barr[i] -= mn
				dfs(aarr, barr, ans+1)
				barr[i] += mn
			}
		}
		aarr[belong_idx] = mn
	} else {
		barr[belong_idx] = 0
		for i := 0; i < len(aarr); i++ {
			if aarr[i] > 0 {
				aarr[i] -= mn
				dfs(aarr, barr, ans+1)
				aarr[i] += mn
			}
		}
		barr[belong_idx] = mn
	}
}
