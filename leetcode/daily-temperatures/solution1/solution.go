// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/description/
package solution

// 单调栈：从右往左
// 时间复杂度：O(n)
// 空间复杂度：O(min(n, U))，其中 U = max(temperatures) - min(temperatures) + 1
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	var st []int
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		temp := temperatures[i]
		for len(st) > 0 && temp >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}
