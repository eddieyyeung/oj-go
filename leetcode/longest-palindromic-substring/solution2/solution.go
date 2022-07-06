package solution2

var maxLen = 0
var maxS = ""
var dp [][]int

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen = 1
	maxS = s[0:1]
	size := len(s)
	dp = make([][]int, size)
	for i, _ := range dp {
		dp[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			fn1(i, j, s)
		}
	}

	//fn2(0, len(s)-1, s)
	return maxS
}

func fn2(left, right int, s string) {
	if right-left+1 <= maxLen {
		return
	}
	if fn1(left, right, s) {
		if right-left+1 > maxLen {
			maxS = s[left : right+1]
			maxLen = right - left + 1
		}
	} else {
		fn2(left+1, right, s)
		fn2(left, right-1, s)
	}
}

func fn1(left, right int, s string) bool {
	if right-left == 1 {
		if s[right] == s[left] {
			dp[left][right] = 1
			if right-left+1 > maxLen {
				maxS = s[left : right+1]
				maxLen = right - left + 1
			}
			return true
		} else {
			dp[left][right] = 2
			return false
		}
	}
	if dp[left][right] == 0 {
		if fn1(left+1, right-1, s) {
			if s[left] == s[right] {
				dp[left][right] = 1
				if right-left+1 > maxLen {
					maxS = s[left : right+1]
					maxLen = right - left + 1
				}
				return true
			} else {
				dp[left][right] = 2
				return false
			}
		} else {
			dp[left][right] = 2
			return false
		}
	} else if dp[left][right] == 1 {
		if right-left+1 > maxLen {
			maxS = s[left : right+1]
			maxLen = right - left + 1
		}
		return true
	} else {
		return false
	}
}
