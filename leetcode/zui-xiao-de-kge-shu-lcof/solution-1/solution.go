package solution_1

import "sort"

// 快排 O(nlogn)
func getLeastNumbers(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}
