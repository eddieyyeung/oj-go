package maximum_swap

// problem: https://leetcode.cn/problems/maximum-swap/

func maximumSwap(num int) (ans int) {
	if num == 0 {
		return 0
	}
	var arr []int
	var max = 0
	var maxIdx = 0
	var a, b int
	var idx int
	for num != 0 {
		r := num % 10
		num = num / 10
		arr = append(arr, r)
		if r < max {
			a = idx
			b = maxIdx
		}
		if r > max {
			maxIdx = idx
			max = r
		}
		idx++
	}
	arr[a], arr[b] = arr[b], arr[a]
	for i := len(arr) - 1; i >= 0; i-- {
		ans = ans*10 + arr[i]
	}
	return
}
