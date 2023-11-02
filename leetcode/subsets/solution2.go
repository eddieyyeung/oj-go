// 78. 子集
// https://leetcode.cn/problems/subsets/description/
package subsets

// 时间复杂度：O(n2^n)
// 其中 n 为 nums 的长度。答案的长度为子集的个数，即 2^n，
// 同时每次递归都把一个数组放入答案，因此会递归 2^n 次，再算上加入答案时需要 O(n) 的时间，
// 所以时间复杂度为 O(n2^n)
//
// 空间复杂度：O(n)O(n)O(n)。返回值的空间不计。
// 方案二：答案的视角
func subsets2(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		// 由于 path 长度可变，需固定答案
		ans = append(ans, append([]int{}, path...))

		// 题目要求不能包含重复的数字
		// 因此规定一个顺序，从当前回溯的 i 至 n 来枚举选择的数字
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			// 恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
