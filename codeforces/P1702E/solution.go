// E. Split Into Two Sets
// https://codeforces.com/problemset/problem/1702/E
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// https://codeforces.com/blog/entry/104763
//  1. Polycarp has n dominoes, on each domino there are 2 numbers.
//     It turns out there will be 2n numbers in total.
//  2. We need to divide 2n numbers (each number from 1 to n) into two sets so that all numbers in each set are different — each set will consist of n numbers.
//     It turns out that all numbers from 1 to n must occur exactly 2 times, no more and no less.
//  3. Let's imagine it all as a bipartite graph, where there are vertices from 1 to n, and dominoes are edges.
//     Since each number occurs exactly 2 times, then we have a lot of cycles.
//     In which the edges of each number must be included in different sets, in other words, the cycles must be of even length.
//  4. This can be checked in O(n) by a simple enumeration.
func solve(ca CaseArg) bool {
	g := make([][]int, ca.N)
	for _, domino := range ca.Dominoes {
		a, b := domino[0]-1, domino[1]-1
		if a == b {
			return false
		}
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		if len(g[a]) > 2 || len(g[b]) > 2 {
			return false
		}
	}

	visited := make([]bool, ca.N)

	for i := 0; i < ca.N; i++ {
		v, cnt := i, 0
		for !visited[v] {
			visited[v] = true
			cnt ^= 1
			for _, nv := range g[v] {
				if !visited[nv] {
					v = nv
					break
				}
			}
		}
		if cnt != 0 {
			return false
		}
	}

	return true
}

type CaseArg struct {
	N        int
	Dominoes [][]int
}

func runCase(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.Dominoes = make([][]int, ca.N)
		for j := 0; j < ca.N; j++ {
			var a, b int
			fmt.Fscan(in, &a)
			fmt.Fscan(in, &b)
			ca.Dominoes[j] = []int{a, b}
		}
		rst := solve(ca)
		if rst {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
