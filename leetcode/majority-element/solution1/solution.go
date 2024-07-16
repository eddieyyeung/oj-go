// https://leetcode.cn/problems/majority-element/description/
package majority_element

// 摩尔投票法, MJRTY - A Fast Majority Vote Algorithm
// https://www.scribd.com/document/319883199/mjrty-ps
func majorityElement(nums []int) int {
	candidate := nums[0]
	cnt := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if candidate == num {
			cnt++
		} else if cnt == 0 {
			candidate = num
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}
