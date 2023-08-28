package insertinterval

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	left, right := newInterval[0], newInterval[1]

	for _, interval := range intervals {
		if interval[0] > right {
			ans = append(ans, []int{left, right})
			left, right = interval[0], interval[1]
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
