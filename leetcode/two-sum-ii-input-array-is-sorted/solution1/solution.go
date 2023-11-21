// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
package solution

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}
