// A. Farmer John's Challenge
// https://codeforces.com/contest/1942/problem/A
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) []int {
	if ca.N == ca.K {
		ret := make([]int, ca.N)
		for i := 0; i < ca.N; i++ {
			ret[i] = 1
		}
		return ret
	} else if ca.K == 1 {
		ret := make([]int, ca.N)
		for i := 0; i < ca.N; i++ {
			ret[i] = i + 1
		}
		return ret
	} else {
		return []int{-1}

	}
}

type CaseArg struct {
	N, K int
}

func runCase(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N, &ca.K)
		ans := solve(ca)
		for _, item := range ans {
			fmt.Fprintf(out, "%d ", item)
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
