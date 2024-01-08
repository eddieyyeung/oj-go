// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/description
package solution

func longestConsecutive(nums []int) int {
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	var ans int = 0
	for k := range mp {
		if !mp[k-1] {
			cnt := 1
			for mp[k+1] {
				cnt++
				k++
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
