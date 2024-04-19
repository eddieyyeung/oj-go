// C. The Sports Festival
// https://codeforces.com/problemset/problem/1509/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func solve(ca CaseArg) int {
	dp := make([][]int, ca.N)
	sort.Ints(ca.Ss)
	for i := 0; i < ca.N; i++ {
		dp[i] = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l >= r {
			return 0
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}
		ret := ca.Ss[r] - ca.Ss[l] + min(dfs(l+1, r), dfs(l, r-1))
		dp[l][r] = ret
		return dp[l][r]
	}

	return dfs(0, ca.N-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type CaseArg struct {
	N  int
	Ss []int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	ca.Ss = make([]int, ca.N)
	for i := 0; i < ca.N; i++ {
		fmt.Fscan(in, &ca.Ss[i])
	}
	ans := solve(ca)
	fmt.Fprintln(out, ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
