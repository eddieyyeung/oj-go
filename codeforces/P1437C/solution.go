// D. Colored Rectangles
// https://codeforces.com/problemset/problem/1398/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if b > a {
		a, b = b, a
	}
	return a - b
}

func solve(ca CaseArg) int {
	sort.Ints(ca.Ts)
	dp := make([][]int, ca.N)
	for i := 0; i < ca.N; i++ {
		dp[i] = make([]int, 401)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(n, t int) int
	dfs = func(n, t int) int {
		if n < 0 {
			return 0
		}
		num := ca.Ts[n]
		if t <= n {
			return 0
		}
		if dp[n][t] != -1 {
			return dp[n][t]
		}
		var ret int = abs(t, num)
		var k int = math.MaxInt32
		for i := n; i <= t-1; i++ {
			d := dfs(n-1, i)
			k = min(k, d)
		}
		ret += k
		dp[n][t] = ret
		return ret
	}
	var ans int = math.MaxInt32

	for i := ca.N; i <= 400; i++ {
		ans = min(ans, dfs(ca.N-1, i))
	}

	return ans
}

type CaseArg struct {
	N  int
	Ts []int
}

func runCase(in io.Reader, out io.Writer) {
	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.Ts = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.Ts[j])
		}
		fmt.Fprintln(out, solve(ca))
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
