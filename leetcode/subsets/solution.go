// 78. 子集
// https://leetcode.cn/problems/subsets/description/
package subsets

// 时间复杂度：O(n2^n)
// 其中 n 为 nums 的长度。每次都是选或不选，递归次数为一个满二叉树的节点个数，那么一共会递归 O(2^n) 次（等比数列和），
// 再算上加入答案时需要 O(n) 的时间，所以时间复杂度为 O(n2^n)
//
// 空间复杂度：O(n)
// 返回值的空间不计。
// 方案一：输入的视角
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			// 由于 path 长度可变，需固定答案
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 1. 子集中不选择 nums[i]
		dfs(i + 1)

		// 2. 子集中选择 nums[i]
		path = append(path, nums[i])
		dfs(i + 1)
		// 恢复现场
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}
