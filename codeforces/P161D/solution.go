// D. Distance in Tree
// https://codeforces.com/problemset/problem/161/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Item struct {
	Next    int
	Visited bool
}

type Graph struct {
	Edges   [][]int
	Visited [][]int
	Size    int
}

func NewGraph(size int) *Graph {
	g := Graph{
		Edges: make([][]int, size),
		Size:  size,
	}
	return &g
}

func (g *Graph) AddEdge(x, y int) {
	g.Edges[x] = append(g.Edges[x], y)
	g.Edges[y] = append(g.Edges[y], x)
}

func solve(ca CaseArg) int {
	g := NewGraph(ca.N)
	for _, edge := range ca.Edges {
		g.AddEdge(edge[0]-1, edge[1]-1)
	}
	dp := make([][]int, ca.N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, ca.K+1)
	}
	var ans int
	var dfs func(parent int, node int)
	dist := make([]int, ca.K+1)
	dfs = func(parent int, node int) {
		dp[node][0] = 1
		for _, child := range g.Edges[node] {
			if child == parent {
				continue
			}
			dfs(node, child)
			for i := 0; i <= ca.K; i++ {
				dist[i] = dp[node][i]
			}
			for i := 0; i < ca.K; i++ {
				ans += dist[ca.K-i-1] * dp[child][i]
				dp[node][i+1] += dp[child][i]
			}
		}
	}
	dfs(0, 0)

	return ans
}

type CaseArg struct {
	N, K  int
	Edges [][2]int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N, &ca.K)
	ca.Edges = make([][2]int, ca.N-1)
	for i := 0; i < ca.N-1; i++ {
		fmt.Fscan(in, &ca.Edges[i][0], &ca.Edges[i][1])
	}
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
