// https://leetcode.cn/problems/majority-element-ii/description/
package majorityelementii

// 博耶-摩尔多数投票算法, MJRTY - A Fast Majority Vote Algorithm
// https://www.scribd.com/document/319883199/mjrty-ps
// https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/
// https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm-for-searching-elements-having-more-than-k-occurrences/
func mjrty(nums []int, k int) []int {
	candidates := make([]int, k)
	votes := make([]int, k)
	for i := 0; i < k; i++ {
		candidates[i] = nums[0]
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		is_vote := false
		for j, candidate := range candidates {
			if candidate == num {
				votes[j]++
				is_vote = true
				break
			}
		}
		if is_vote {
			continue
		}
		is_change_candidate := false
		for j, vote := range votes {
			if vote == 0 {
				candidates[j] = num
				votes[j]++
				is_change_candidate = true
				break
			}
		}
		if is_change_candidate {
			continue
		}
		for j := range votes {
			votes[j]--
		}
	}
	m := make(map[int]bool)
	var ret []int
	for _, candidate := range candidates {
		if m[candidate] {
			continue
		}
		cnt := 0
		for _, num := range nums {
			if candidate == num {
				cnt++
			}
		}
		if cnt > len(nums)/k {
			ret = append(ret, candidate)
			m[candidate] = true
		}
	}
	return ret
}

func majorityElement(nums []int) []int {
	return mjrty(nums, 3)
}
