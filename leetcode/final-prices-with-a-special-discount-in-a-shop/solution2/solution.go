package solution2

// 方法二：单调栈
func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		vi := prices[i]
		sti := len(st) - 1
		for j := len(st) - 1; j >= 0; j-- {
			if st[j] > vi {
				sti--
			}
		}
		ans[i] = vi - st[sti]
		st = st[:sti+1]
		st = append(st, vi)
	}
	return ans
}
