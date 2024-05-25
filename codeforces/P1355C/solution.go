// C. Count Triangles
// https://codeforces.com/problemset/problem/1355/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) int {
	var sum = [1e6 + 5]int{}
	// sum[i+B]+1, sum[i+C+1]-1
	for i := ca.A; i <= ca.B; i++ {
		sum[i+ca.B]++
		sum[i+ca.C+1]--
	}
	// 前缀和，计算出每个 s 出现的组合个数
	for i := ca.A + ca.B; i <= ca.B+ca.C+1; i++ {
		sum[i] += sum[i-1]
	}

	// 后缀和，累加 cnt_s
	for i := ca.B + ca.C; i >= ca.C; i-- {
		sum[i] += sum[i+1]
	}

	// 遍历 z [C,D]，累加每一个小于 s 的结果 sum[z_i+1]，求得结果
	var ans int
	for i := ca.C; i <= ca.D; i++ {
		ans += sum[i+1]
	}

	return ans
}

type CaseArg struct {
	A, B, C, D int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.A, &ca.B, &ca.C, &ca.D)
	ans := solve(ca)
	fmt.Fprintln(out, ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
