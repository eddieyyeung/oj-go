// 2398. 预算内的最多机器人数目
// https://leetcode.cn/problems/maximum-number-of-robots-within-budget/description/
package solution

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	var ans int
	n := len(chargeTimes)
	var st []int
	var sum int
	var j int
	for i := 0; i < n; i++ {
		// in
		for len(st) > 0 && chargeTimes[i] >= chargeTimes[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
		sum += runningCosts[i]
		// out
		for len(st) > 0 {
			if chargeTimes[st[0]]+(i-j+1)*sum > int(budget) {
				if st[0] == j {
					st = st[1:]
				}
				sum -= runningCosts[j]
				j++
			} else {
				break
			}
		}
		// ans
		ans = max(ans, i-j+1)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
