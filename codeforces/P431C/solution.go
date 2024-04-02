// C. k-Tree
// https://codeforces.com/problemset/problem/431/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const mod = 1e9 + 7

func solve(ca CaseArg) int {
	dp := make([][2]int, ca.N+1)
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 1
	var dfs func(n, s int) int
	dfs = func(n, s int) int {
		if s == 1 && n < ca.D {
			return 0
		}
		if n < 0 {
			return 0
		}
		if dp[n][s] != -1 {
			return dp[n][s]
		}
		var rst int
		if s == 0 {
			for i := 1; i <= ca.D-1; i++ {
				rst += dfs(n-i, 0)
			}
		} else {
			for i := ca.D; i <= ca.K; i++ {
				rst += dfs(n-i, 0)
			}
			for i := 1; i <= ca.K; i++ {
				rst += dfs(n-i, 1)
			}
		}
		rst = rst % mod
		dp[n][s] = rst
		return rst
	}
	ans := dfs(ca.N, 1)
	return ans
}

type CaseArg struct {
	N, K, D int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	fmt.Fscan(in, &ca.K)
	fmt.Fscan(in, &ca.D)
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
