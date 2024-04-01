// A. Timofey and a tree
// https://codeforces.com/problemset/problem/763/A
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) int {
	diff := make([]int, ca.N)
	cnt := 0
	for _, edge := range ca.Edges {
		if v, w := edge[0]-1, edge[1]-1; ca.Colors[v] != ca.Colors[w] {
			diff[v]++
			diff[w]++
			cnt++
		}
	}
	for i, n := range diff {
		if n == cnt {
			return i + 1
		}
	}
	return -1
}

// func solve(ca CaseArg) int {
// 	g := make([][]int, ca.N+1)
// 	for _, edge := range ca.Edges {
// 		a, b := edge[0], edge[1]
// 		g[a] = append(g[a], b)
// 		g[b] = append(g[b], a)
// 	}
// 	var root1, root2 int
// 	for i := 1; i <= ca.N; i++ {
// 		for _, v := range g[i] {
// 			if ca.Colors[i-1] != ca.Colors[v-1] {
// 				root1 = i
// 				root2 = v
// 			}
// 		}
// 	}
// 	if root1 == 0 {
// 		return 1
// 	}
// 	visited := make([]int, ca.N+1)
// 	var dfs func(i int, father int) bool
// 	dfs = func(i int, father int) bool {
// 		if visited[i] != 0 {
// 			return visited[i] == 1
// 		}
// 		var rst int = 1
// 		for _, v := range g[i] {
// 			if v == father {
// 				continue
// 			}
// 			if ca.Colors[v-1] != ca.Colors[i-1] || !dfs(v, i) {
// 				rst = 2
// 				break
// 			}
// 		}
// 		visited[i] = rst
// 		return visited[i] == 1
// 	}

// 	isFound := func(root int) bool {
// 		is_found := true
// 		for _, v := range g[root] {
// 			visited[v] = 0
// 			if !dfs(v, root) {
// 				is_found = false
// 				break
// 			}
// 		}
// 		return is_found
// 	}
// 	if isFound(root1) {
// 		return root1
// 	}
// 	if isFound(root2) {
// 		return root2
// 	}
// 	return -1
// }

type CaseArg struct {
	N      int
	Edges  [][2]int
	Colors []int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	ca.Edges = make([][2]int, ca.N-1)
	for i := 0; i < ca.N-1; i++ {
		fmt.Fscan(in, &ca.Edges[i][0])
		fmt.Fscan(in, &ca.Edges[i][1])
	}
	ca.Colors = make([]int, ca.N)
	for i := 0; i < ca.N; i++ {
		fmt.Fscan(in, &ca.Colors[i])
	}
	rst := solve(ca)
	if rst == -1 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, rst)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
