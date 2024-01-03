// 238. 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/description
package solution

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	product := 1
	for i := 0; i < n; i++ {
		ans[i] = product
		product *= nums[i]
	}
	product = 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= product
		product *= nums[i]
	}
	return ans
}
