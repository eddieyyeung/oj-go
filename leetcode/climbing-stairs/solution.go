// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description/
package climbingstairs

func climbStairs(n int) int {
	dp0, dp1 := 1, 1
	for i := 2; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}
