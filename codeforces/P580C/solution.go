// C. Kefa and Park
// https://codeforces.com/problemset/problem/580/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func runCase(ca CaseArg) {
	rst := solve(ca)
	ptext := fmt.Sprintf("%v", rst)
	fmt.Println(ptext)
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanCase(r io.Reader) CaseArg {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var ca CaseArg
	ca.N = scanInt(scanner)
	ca.M = scanInt(scanner)
	ca.CatVertices = make([]int, ca.N)
	ca.Vertices = make([][]int, ca.N)
	for i := 0; i < ca.N; i++ {
		ca.CatVertices[i] = scanInt(scanner)
	}
	for i := 0; i < ca.N-1; i++ {
		x, y := scanInt(scanner), scanInt(scanner)
		if x > y {
			x, y = y, x
		}
		ca.Vertices[x-1] = append(ca.Vertices[x-1], y-1)
		ca.Vertices[y-1] = append(ca.Vertices[y-1], x-1)
	}
	return ca
}

type CaseArg struct {
	N           int // the number of vertices of the tree
	M           int // the maximum number of consecutive vertices with cats that is still ok for Kefa
	CatVertices []int
	Vertices    [][]int
}

func main() {
	runCase(scanCase(os.Stdin))
}
