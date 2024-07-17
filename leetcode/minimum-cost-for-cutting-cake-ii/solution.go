// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-ii/
package minimumcostforcuttingcakeii

import "sort"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	sort.Slice(horizontalCut, func(i, j int) bool {
		return horizontalCut[i] > horizontalCut[j]
	})
	sort.Slice(verticalCut, func(i, j int) bool {
		return verticalCut[i] > verticalCut[j]
	})
	var cnt_h int = 0
	var cnt_v int = 0
	var ans int64 = 0
	for cnt_h < m-1 && cnt_v < n-1 {
		if horizontalCut[cnt_h] > verticalCut[cnt_v] {
			ans += int64(horizontalCut[cnt_h] * (cnt_v + 1))
			cnt_h++
		} else {
			ans += int64(verticalCut[cnt_v] * (cnt_h + 1))
			cnt_v++
		}
	}
	for cnt_h < m-1 {
		ans += int64(horizontalCut[cnt_h] * (cnt_v + 1))
		cnt_h++
	}
	for cnt_v < n-1 {
		ans += int64(verticalCut[cnt_v] * (cnt_h + 1))
		cnt_v++
	}
	return ans
}
