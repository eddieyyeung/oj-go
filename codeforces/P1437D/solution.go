// D. Minimal Height Tree
// https://codeforces.com/problemset/problem/1437/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CaseArg struct {
	N        int
	Vertices []int
}

func solve(ca CaseArg) int {
	if ca.N == 1 {
		return 0
	}
	var i int = 1
	var j int = 0
	var ans int = 0
	for i < ca.N {
		ans++
		t := i
		for j < t {
			for i+1 < ca.N && ca.Vertices[i] < ca.Vertices[i+1] {
				i++
			}
			j++
			i++
		}
	}
	return ans
}

func runCase(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.Vertices = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			var vertex int
			fmt.Fscan(in, &vertex)
			ca.Vertices[j] = vertex
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
