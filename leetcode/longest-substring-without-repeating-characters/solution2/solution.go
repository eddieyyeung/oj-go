package solution2

// 同向双指针 滑动窗口
func lengthOfLongestSubstring(s string) int {
	haschar := make([]int, 256)
	ans := 0
	left := 0
	for right, ch := range s {
		idx := ch - ' '
		haschar[idx]++
		for left < right && haschar[idx] > 1 {
			haschar[s[left]-' ']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
