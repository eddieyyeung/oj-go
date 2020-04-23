// Package solution2 ...
// https://leetcode-cn.com/problems/longest-consecutive-sequence/
package solution2

func longestConsecutive(nums []int) int {
	var numsMap = map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	longest := 0
	for k, _ := range numsMap {
		if _, existed := numsMap[k-1]; !existed {
			current := k + 1
			l := 1
			for {
				if _, existed := numsMap[current]; existed {
					l++
					current++
				} else {
					break
				}
			}
			if l > longest {
				longest = l
			}
		}
	}
	return longest
}
