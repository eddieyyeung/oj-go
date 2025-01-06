// C. Engineer Artem
// https://codeforces.com/problemset/problem/1438/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CaseArg struct {
	N      int
	M      int
	Matrix [][]int
}

func solve(ca CaseArg) [][]int {
	for i := 0; i < ca.N; i++ {
		for j := 0; j < ca.M; j++ {
			if (i+j)&1 == 0 {
				ca.Matrix[i][j] += ca.Matrix[i][j] & 1
			} else {
				ca.Matrix[i][j] += 1 - ca.Matrix[i][j]&1
			}
		}
	}
	return ca.Matrix
}
func runCase(in io.Reader, out io.Writer) {
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N, &ca.M)
		ca.Matrix = make([][]int, ca.N)
		for j := 0; j < ca.N; j++ {
			ca.Matrix[j] = make([]int, ca.M)
		}
		for j := 0; j < ca.N; j++ {
			for k := 0; k < ca.M; k++ {
				fmt.Fscan(in, &ca.Matrix[j][k])
			}
		}
		rst := solve(ca)
		for j := 0; j < ca.N; j++ {
			for k := 0; k < ca.M; k++ {
				fmt.Fprintf(out, "%d ", rst[j][k])
			}
			fmt.Fprintln(out)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
