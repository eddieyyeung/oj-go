package mean_of_array_after_removing_some_elements

import (
	"sort"
)

// problem: https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	sum := 0.0
	arr = arr[len(arr)/20 : len(arr)-(len(arr)/20)]
	for _, v := range arr {
		sum += float64(v)
	}
	return sum / float64(len(arr))
}
