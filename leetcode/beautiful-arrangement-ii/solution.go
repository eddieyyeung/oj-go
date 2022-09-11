package beautiful_arrangement_ii

// problem: https://leetcode.cn/problems/beautiful-arrangement-ii/

func constructArray(n int, k int) []int {
	if k > n-1 {
		return nil
	}
	ans := make([]int, 0, n)
	a := 1
	b := k + 1
	idx := 0
	for i := 0; i <= (k+1)/2; i++ {
		if a == b {
			ans = append(ans, a)
			idx++
		} else if a > b {
			break
		} else {
			ans = append(ans, a, b)
			a++
			b--
			idx += 2
		}
	}
	for i := idx; i < n; i++ {
		ans = append(ans, i+1)
	}
	return ans
}
