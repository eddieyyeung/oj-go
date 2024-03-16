// C. Kefa and Park
// https://codeforces.com/problemset/problem/580/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CaseArg struct {
	N           int // the number of vertices of the tree
	M           int // the maximum number of consecutive vertices with cats that is still ok for Kefa
	CatVertices []int
	Vertices    [][]int
}

func solve(ca CaseArg) int {
	var ans int
	visited := make([]int, ca.N)
	var dfs func(node int, cat int)
	dfs = func(node, cat int) {
		if ca.CatVertices[node] == 1 {
			cat += 1
		} else {
			cat = 0
		}
		if cat > ca.M {
			return
		}
		isLeaf := true
		for _, vertex := range ca.Vertices[node] {
			if visited[vertex] == 0 {
				isLeaf = false
				visited[vertex] = 1
				dfs(vertex, cat)
				visited[vertex] = 0
			}
		}
		if isLeaf {
			ans++
		}
	}
	visited[0] = 1
	dfs(0, 0)
	return ans
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	fmt.Fscan(in, &ca.M)
	ca.CatVertices = make([]int, ca.N)
	ca.Vertices = make([][]int, ca.N)
	for i := 0; i < ca.N; i++ {
		fmt.Fscan(in, &ca.CatVertices[i])
	}
	for i := 0; i < ca.N-1; i++ {
		var x, y int
		fmt.Fscan(in, &x)
		fmt.Fscan(in, &y)
		if x > y {
			x, y = y, x
		}
		ca.Vertices[x-1] = append(ca.Vertices[x-1], y-1)
		ca.Vertices[y-1] = append(ca.Vertices[y-1], x-1)
	}
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
