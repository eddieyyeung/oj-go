package solution1

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	min := math.MaxInt
	count := 0
	for i := len(pairs) - 1; i >= 0; i-- {
		if pairs[i][1] < min {
			count++
			min = pairs[i][0]
			continue
		}
	}
	return count
}
