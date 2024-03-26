// D. Min Cost String
// https://codeforces.com/problemset/problem/1511/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// dfs
func solve2(ca CaseArg) string {
	cur := make([]int, ca.K)
	var path []int
	var dfs func(u int)

	dfs = func(u int) {
		for cur[u] < ca.K {
			v := cur[u]
			cur[u]++
			dfs(v)
			path = append(path, v)
		}
	}
	dfs(0)
	var ans strings.Builder
	ans.WriteByte('a')
	for i := 0; i < ca.N-1; i++ {
		ans.WriteByte(byte(path[i%len(path)] + 'a'))
	}
	return ans.String()
}

// bfs
func solve(ca CaseArg) string {
	visited := make([]int, ca.K)
	var path []int = []int{}
	q := []int{0}
	var curr int = 0
	visited[0] = 1
	for len(q) > 0 {
		if visited[curr] < ca.K {
			q = append(q, curr)
			next := visited[curr]
			visited[curr]++
			curr = next
		} else {
			path = append(path, curr)
			curr = q[len(q)-1]
			q = q[:len(q)-1]
		}
	}

	var ans strings.Builder
	ans.WriteByte(byte(0 + 'a'))
	for i := 0; i < ca.N-1; i++ {
		ans.WriteByte(byte(path[i%len(path)] + 'a'))
	}
	return ans.String()
}

type CaseArg struct {
	N int
	K int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	fmt.Fscan(in, &ca.K)
	rst := solve(ca)
	fmt.Fprintln(out, rst)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
