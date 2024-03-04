// 887. 鸡蛋掉落
// https://leetcode.cn/problems/super-egg-drop/description/
package solution

// k eggs, 1-n floors
// f 1 <= f <= n, ny egg dropped at a floor higher than f will break
// return mininum number of moves

// dp(k, n) = 1 + min{1<=x<=n}(max(dp(k, n-x), dp(k-1,x-1)))
func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	var dp func(k, n int) int
	dp = func(k, n int) int {
		if n == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		if memo[k][n] != -1 {
			return memo[k][n]
		}
		l, r := 1, n
		for r-l > 1 {
			m := (l + r) / 2
			c1 := dp(k-1, m-1)
			c2 := dp(k, n-m)
			if c1 >= c2 {
				r = m
			} else {
				l = m
			}
		}
		ans := min(max(dp(k-1, l-1), dp(k, n-l)), max(dp(k-1, r-1), dp(k, n-r))) + 1

		memo[k][n] = ans
		return ans
	}

	return dp(k, n)
}

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
