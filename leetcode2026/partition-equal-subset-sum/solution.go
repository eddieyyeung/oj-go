package partition_equal_subset_sum

// 示例 1：
//
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 示例 2：
//
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
//
// 提示：
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	t := sum / 2
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, t+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(i int, c int) bool
	dfs = func(i int, c int) bool {
		if c == 0 {
			return true
		}
		if c < 0 {
			return false
		}
		if i == n {
			return false
		}
		if dp[i][c] != -1 {
			return dp[i][c] == 1
		}
		if dfs(i+1, c) || dfs(i+1, c-nums[i]) {
			dp[i][c] = 1
		} else {
			dp[i][c] = 0
		}
		return dp[i][c] == 1
	}
	return dfs(0, t)
}
