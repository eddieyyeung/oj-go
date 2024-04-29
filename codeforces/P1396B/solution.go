// D. Carousel
// https://codeforces.com/problemset/problem/1328/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(ca CaseArg) string {
	var m int
	var s int
	for _, a := range ca.As {
		m = max(m, a)
		s += a
	}
	if s&1 == 1 || m*2 > s {
		return "T"
	}
	return "HL"
}

type CaseArg struct {
	N  int
	As []int
}

func runCase(in io.Reader, out io.Writer) {
	var case_num int
	fmt.Fscan(in, &case_num)
	for i := 0; i < case_num; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.As = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.As[j])
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
