// 1499. 满足不等式的最大值
// https://leetcode.cn/problems/max-value-of-equation/description/
package solution

func findMaxValueOfEquation(points [][]int, k int) int {
	n := len(points)
	var q []int
	var ans int = -1e9
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		for len(q) > 0 && x-points[q[0]][0] > k {
			q = q[1:]
		}
		if len(q) > 0 {
			x_q0, y_q0 := points[q[0]][0], points[q[0]][1]
			e := x - x_q0 + y_q0 + y
			ans = max(ans, e)
		}
		sub := x - y
		for len(q) > 0 {
			x_q, y_q := points[q[len(q)-1]][0], points[q[len(q)-1]][1]
			subq := x_q - y_q
			if sub > subq {
				break
			}
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
