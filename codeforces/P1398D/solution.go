// D. Colored Rectangles
// https://codeforces.com/problemset/problem/1398/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(ca CaseArg) int {
	for i := 0; i < 3; i++ {
		sort.Ints(ca.Sticks[i])
	}
	dp := make([][][]int, ca.Nums[0]+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, ca.Nums[1]+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, ca.Nums[2]+1)
			for k := 0; k < len(dp[i][j]); k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, k int) int

	dfs = func(i, j, k int) int {
		if dp[i][j][k] != -1 {
			return dp[i][j][k]
		}
		var c1, c2, c3 int
		if i != 0 && j != 0 {
			c1 = ca.Sticks[0][i-1]*ca.Sticks[1][j-1] + dfs(i-1, j-1, k)
		}
		if i != 0 && k != 0 {
			c2 = ca.Sticks[0][i-1]*ca.Sticks[2][k-1] + dfs(i-1, j, k-1)
		}
		if j != 0 && k != 0 {
			c3 = ca.Sticks[1][j-1]*ca.Sticks[2][k-1] + dfs(i, j-1, k-1)
		}
		ret := max(max(c1, c2), c3)
		dp[i][j][k] = ret
		return dp[i][j][k]
	}
	var ans int = dfs(ca.Nums[0], ca.Nums[1], ca.Nums[2])
	return ans
}

type CaseArg struct {
	Nums   [3]int
	Sticks [3][]int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.Nums[0], &ca.Nums[1], &ca.Nums[2])
	for i := 0; i < 3; i++ {
		ca.Sticks[i] = make([]int, ca.Nums[i])
		for j := 0; j < ca.Nums[i]; j++ {
			fmt.Fscan(in, &ca.Sticks[i][j])
		}
	}
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
