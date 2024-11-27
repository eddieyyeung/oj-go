// K. Grid Walk
// https://codeforces.com/problemset/problem/2038/K
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Item struct {
	A int
	B int
}

func gcd(a, b int) int {
	var ret int
	if b == 0 {
		ret = a
	} else {
		ret = gcd(b, a%b)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(ca CaseArg) int {
	dp := make([]int, 120)
	max_y := ca.N - 1
	max_x := ca.N - 1
	for gcd(max_x, ca.A) != 1 {
		max_x--
	}
	for gcd(max_y, ca.B) != 1 {
		max_y--
	}
	var ans int = max_x - 1 + max_y - 1
	for i := 1; i <= max_x; i++ {
		ans += gcd(i, ca.A)
	}
	for j := 1; j <= max_y; j++ {
		ans += gcd(j, ca.B)
	}
	dp[0] = ans
	for i := 1; i < ca.N-max_y+1; i++ {
		dp[i] = dp[i-1] + gcd(ca.A, max_x) + gcd(ca.B, max_y+i)
	}
	for i := 1; i < ca.N-max_x+1; i++ {
		for j := 0; j < ca.N-max_y+1; j++ {
			gdcsum := gcd(ca.A, max_x+i) + gcd(ca.B, max_y+j)
			c := dp[j] + gdcsum
			if j != 0 {
				c = min(c, dp[j-1]+gdcsum)
			}
			dp[j] = c
			ans = c
		}
	}
	return ans
}

type CaseArg struct {
	N, A, B int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N, &ca.A, &ca.B)
	ans := solve(ca)
	fmt.Fprintln(out, ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
