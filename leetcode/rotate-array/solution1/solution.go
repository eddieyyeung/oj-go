// 189. 轮转数组
// https://leetcode.cn/problems/rotate-array/description
package solution

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	reverse(nums)
	reverse(nums[:(k % n)])
	reverse(nums[(k % n):])
}
