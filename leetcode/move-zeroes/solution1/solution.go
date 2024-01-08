// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/description
package solution

func moveZeroes(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
