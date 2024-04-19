// D. Multiplication Table
// https://codeforces.com/problemset/problem/448/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(ca CaseArg) int {
	low := 1
	high := ca.M * ca.N
	get_count := func(num int) int {
		count := 0
		for i := 0; i < ca.N; i++ {
			count += min(num/(i+1), ca.M)
		}
		return count
	}

	for low <= high {
		mid := (low + high) / 2
		mid_count := get_count(mid)
		if mid_count >= ca.K {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// fmt.Println(low, high)
	return low
}

type CaseArg struct {
	N, M, K int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N, &ca.M, &ca.K)
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
