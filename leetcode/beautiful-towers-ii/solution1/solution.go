// 2866. 美丽塔 II
// https://leetcode.cn/problems/beautiful-towers-ii/description/
package solution

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	presum := make([]int, n+1)
	var sum int = 0
	q := []int{-1}
	for i := 0; i < n; i++ {
		h := maxHeights[i]
		for len(q) > 1 && maxHeights[q[len(q)-1]] >= h {
			t := q[len(q)-1]
			q = q[:len(q)-1]
			sum -= maxHeights[t] * (t - q[len(q)-1])
		}
		sum += (i - q[len(q)-1]) * h
		presum[i] = sum
		q = append(q, i)
	}
	q = []int{n}
	sum = 0
	var ans int
	for i := n - 1; i >= 0; i-- {
		h := maxHeights[i]
		for len(q) > 1 && maxHeights[q[len(q)-1]] >= h {
			t := q[len(q)-1]
			q = q[:len(q)-1]
			sum -= maxHeights[t] * (q[len(q)-1] - t)
		}
		sum += (q[len(q)-1] - i) * h
		ans = max(ans, sum+presum[i]-h)
		q = append(q, i)
	}

	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
