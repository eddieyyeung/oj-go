// https://leetcode.cn/problems/gas-station
package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			idx := (i + j) % n
			sum += gas[idx] - cost[idx]
			if sum < 0 {
				i = i + j
				break
			}
		}
		if sum >= 0 {
			return i
		}
	}
	return -1
}
