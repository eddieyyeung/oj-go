// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/description/
package solution

// 单调栈：从左往右
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func dailyTemperatures(temperatures []int) []int {
	var st []int
	ans := make([]int, len(temperatures))
	for index, temp := range temperatures {
		n := len(st)
		for i := n - 1; i >= 0; i-- {
			if temp > temperatures[st[i]] {
				ans[st[i]] = index - st[i]
				st = st[:len(st)-1]
			} else {
				break
			}
		}
		st = append(st, index)
	}
	return ans
}
