package applyoperationstomaketwostringsequal

import (
	"strings"
)

func minOperations(s1 string, s2 string, x int) int {
	if strings.Count(s1, "1")%2 != strings.Count(s2, "1")%2 {
		return -1
	}
	n := len(s1)
	memo := make([][][2]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(i int, rcnt int, pre int) int

	// i: 当前下标
	// rcnt: 可使用的免费反转次数，由操作 1 产生
	// pre: 之前的数是否进行过反转，由操作 2 产生
	dfs = func(i, rcnt int, pre int) int {
		if i < 0 {
			if rcnt > 0 || pre != 0 {
				return 1e9
			}
			return 0
		}
		if p := memo[i][rcnt][pre]; p != -1 {
			return p
		}
		// 1. 如果之前已经进行操作 2，且当前位置字符不相等，则不用反转
		// 2. 如果当前位置已经向灯，则不用反转
		if (s1[i] == s2[i]) == (pre == 0) {
			return dfs(i-1, rcnt, 0)
		}

		// case1 使用操作1，并新增一次可免费使用的反转
		ans1 := dfs(i-1, rcnt+1, 0) + x
		// case2 使用操作2，并
		ans2 := dfs(i-1, rcnt, 1) + 1
		ans := min(ans1, ans2)
		// case3 使用操作1，并使用一次免费可用的反转
		if rcnt > 0 {
			ans = min(ans, dfs(i-1, rcnt-1, 0))
		}
		memo[i][rcnt][pre] = ans
		return ans
	}
	return dfs(len(s1)-1, 0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
