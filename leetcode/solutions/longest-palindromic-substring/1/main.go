package solution

var count int = 0

func recur(dp *[1000][1000]int, i int, j int, result *string, s string) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	if i == j {
		dp[i][j] = 1
		return dp[i][j]
	}
	if j-i == 1 {
		if s[j] == s[i] {
			dp[i][j] = 1
			if 2 > len(*result) {
				*result = s[i : j+1]
			}
		} else {
			dp[i][j] = 2
		}
		return dp[i][j]
	}
	count++

	// fmt.Println(i, j)
	var p int
	if dp[i+1][j-1] == 0 {
		p = recur(dp, i+1, j-1, result, s)
	} else {
		p = dp[i+1][j-1]
	}
	if p == 1 && s[i] == s[j] {
		dp[i][j] = 1
		if v := j - i + 1; v > len(*result) {
			*result = s[i : j+1]
		}
	} else {
		dp[i][j] = 2
	}
	return dp[i][j]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := [1000][1000]int{}
	result := s[0:1]
	for i := 0; i < len(s)-len(result); i++ {
		for j := i + len(result); j < len(s); j++ {
			recur(&dp, i, j, &result, s)
		}
	}
	// fmt.Println("len", len(s))
	// fmt.Println("count", count)
	return result
}

func Run(s string) string {
	return longestPalindrome(s)
}
