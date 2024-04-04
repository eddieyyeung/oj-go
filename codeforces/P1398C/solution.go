// C. Good Subarrays
// https://codeforces.com/problemset/problem/1398/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) int {
	var sum int = 0
	m := map[int]int{0: 1}
	var ans int
	for i := 1; i <= ca.N; i++ {
		sum += int(ca.A[i-1] & 15)
		ans += m[sum-i]
		m[sum-i]++
	}
	return ans
}

type CaseArg struct {
	N int
	A []byte
}

func runCase(in io.Reader, out io.Writer) {
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		fmt.Fscan(in, &ca.A)
		ans := solve(ca)
		fmt.Fprintln(out, ans)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
