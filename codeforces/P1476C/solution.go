// C. Longest Simple Cycle
// https://codeforces.com/problemset/problem/1476/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) int {
	dp := make([]int, ca.N)
	dp[0] = abs(ca.Bs[1] - ca.As[1])
	var ans int
	for i := 1; i < ca.N; i++ {
		if ca.As[i] == ca.Bs[i] {
			dp[i] = ca.Cs[i] - 1 + 2
		} else {
			diff := abs(ca.As[i] - ca.Bs[i])
			dp[i] = max(dp[i-1]-diff+(ca.Cs[i]-1+2), diff+(ca.Cs[i]-1+2))
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type CaseArg struct {
	// the number of chains you have.
	N  int
	Cs []int
	As []int
	Bs []int
}

func runCase(in io.Reader, out io.Writer) {
	// t cases
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.Cs = make([]int, ca.N)
		ca.As = make([]int, ca.N)
		ca.Bs = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			var n int
			fmt.Fscan(in, &n)
			ca.Cs[j] = n
		}
		for j := 0; j < ca.N; j++ {
			var n int
			fmt.Fscan(in, &n)
			ca.As[j] = n
		}
		for j := 0; j < ca.N; j++ {
			var n int
			fmt.Fscan(in, &n)
			ca.Bs[j] = n
		}
		rst := solve(ca)
		fmt.Fprintln(out, rst)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
