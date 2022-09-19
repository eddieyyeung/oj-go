package sort_array_by_increasing_frequency

import (
	"sort"
)

// problem: https://leetcode.cn/problems/sort-array-by-increasing-frequency/

type Item struct {
	Num   int
	Count int
}

func frequencySort(nums []int) []int {
	idxes := make([]Item, 201)
	for _, num := range nums {
		idxes[num+100].Num = num
		idxes[num+100].Count++
	}
	sort.Slice(idxes, func(i, j int) bool {
		if idxes[i].Count > idxes[j].Count {
			return false
		} else if idxes[i].Count < idxes[j].Count {
			return true
		} else {
			return idxes[i].Num > idxes[j].Num
		}
	})
	ans := make([]int, 0, len(nums))
	for _, idx := range idxes {
		for i := 0; i < idx.Count; i++ {
			ans = append(ans, idx.Num)
		}
	}
	return ans
}
