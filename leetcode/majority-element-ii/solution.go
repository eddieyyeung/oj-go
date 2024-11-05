// https://leetcode.cn/problems/majority-element-ii/description/
package majorityelementii

// 博耶-摩尔多数投票算法, MJRTY - A Fast Majority Vote Algorithm
// https://www.scribd.com/document/319883199/mjrty-ps
// https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/
// https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm-for-searching-elements-having-more-than-k-occurrences/
func mjrty(nums []int, k int) []int {
	var INF int = 1e9 + 1
	candidates := make([][2]int, k-1)
	for i := 0; i < len(candidates); i++ {
		candidates[i][0] = INF
	}
	var ret []int
	for _, num := range nums {
		is_found := false
		for j := 0; j < k-1; j++ {
			if candidates[j][0] == num {
				candidates[j][1]++
				is_found = true
				break
			}
		}
		if !is_found {
			for j := 0; j < k-1; j++ {
				if candidates[j][1] == 0 {
					candidates[j][0] = num
					candidates[j][1] = 1
					is_found = true
					break
				}
			}
		}
		if !is_found {
			for j := 0; j < k-1; j++ {
				candidates[j][1]--
			}
		}
	}
	for _, candidate := range candidates {
		if candidate[0] == INF {
			continue
		}
		cnt := 0
		for _, num := range nums {
			if candidate[0] == num {
				cnt++
			}
		}
		if cnt > len(nums)/k {
			ret = append(ret, candidate[0])
		}
	}
	return ret
}

func majorityElement(nums []int) []int {
	return mjrty(nums, 3)
}
