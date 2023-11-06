// 46. 全排列
// https://leetcode.cn/problems/permutations/description/
package permutations

// 时间复杂度：O(n * n!)，其中 n 为 nums 的长度。
// 视频中提到，搜索树中的节点个数低于 3⋅n!。实际上，精确值为 ⌊e⋅n!⌋，其中 e=2.718... 为自然常数。
// 每个非叶节点要花费 O(n) 的时间遍历 onPath 数组，每个叶结点也要花费 O(n) 的时间复制 path 数组，因此时间复杂度为 O(n⋅n!)。
//
// 空间复杂度：O(n)。返回值的空间不计入。
func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	path := make([]int, n)
	// 当前未选数字集合
	onPath := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
