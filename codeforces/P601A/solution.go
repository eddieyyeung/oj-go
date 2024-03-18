// A. The Two Routes
// https://codeforces.com/problemset/problem/601/A
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

// n towns
// m bidirectionay railways
// road: for each pair of different towns x and y,
//       there is a bidirectional road between towns x and y if and only if there is no railway between them
// Travelling to a different town using one railway or one road always takes exactly one hour.

// 1. A train and a bus leave town 1 at the same time. They both have the same destination, town n.
// 2. The train can move only along railways and the bus can move only along roads.
// 3. One of the most important aspects to consider is safety — in order to avoid accidents at railway crossings,
//    the train and the bus must not arrive at the same town (except town n) simultaneously.

// ask:
// what is the minimum number of hours needed for both vehicles to reach town n (the maximum of arrival times of the bus and the train)?

// input
// n, the number of towns, 2 ≤ n ≤ 400
// m, the number of railways, 0 ≤ m ≤ n(n - 1) / 2
// Each of the next m lines contains two integers u and v, denoting a railway between towns u and v (1 ≤ u, v ≤ n, u ≠ v).
// You may assume that there is at most one railway connecting any two towns.

// output
// Output one integer — the smallest possible time of the later vehicle's arrival in town n.
// If it's impossible for at least one of the vehicles to reach town n, output  - 1.

func solve(ca CaseArg) int {
	maxN := math.MaxInt
	g := make([][]int, ca.N)
	for i := 0; i < ca.N; i++ {
		g[i] = make([]int, ca.N)
	}
	for _, r := range ca.Railways {
		u, v := r[0]-1, r[1]-1
		g[u][v] = 1
		g[v][u] = 1
	}
	visited_railway := make([]bool, ca.N)
	visited_road := make([]bool, ca.N)
	path_railway := make([]int, ca.N)
	path_road := make([]int, ca.N)
	for i := 1; i < ca.N; i++ {
		path_railway[i] = maxN
		path_road[i] = maxN
	}
	q := []int{0}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		if visited_railway[u] {
			continue
		}
		visited_railway[u] = true
		for v := 0; v < len(g[u]); v++ {
			if !visited_railway[v] && g[u][v] == 1 {
				path_railway[v] = min(path_railway[v], path_railway[u]+1)
				q = append(q, v)
			}
		}
	}

	q = []int{0}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		if visited_road[u] {
			continue
		}
		visited_road[u] = true
		for v := 0; v < len(g[u]); v++ {
			if !visited_road[v] && g[u][v] == 0 {
				path_road[v] = min(path_road[v], path_road[u]+1)
				if path_road[v] == path_railway[v] {
					path_road[v]++
				}
				q = append(q, v)
			}
		}
	}

	ans := max(path_railway[ca.N-1], path_road[ca.N-1])
	if ans >= maxN {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type CaseArg struct {
	N        int
	M        int
	Railways [][]int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	fmt.Fscan(in, &ca.M)
	for i := 0; i < ca.M; i++ {
		var u, v int
		fmt.Fscan(in, &u)
		fmt.Fscan(in, &v)
		ca.Railways = append(ca.Railways, []int{u, v})
	}
	rst := solve(ca)
	fmt.Fprintln(out, rst)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
