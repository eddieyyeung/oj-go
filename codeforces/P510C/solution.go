// C. Fox And Names
// https://codeforces.com/problemset/problem/510/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Graph struct {
	Size     int
	Edges    [][]int
	InDegree []int
}

func (graph *Graph) AddEdge(from, to int) {
	graph.Edges[from] = append(graph.Edges[from], to)
	graph.InDegree[to]++
}

func (graph *Graph) TopSort() ([]int, bool) {
	queue := []int{}
	for i := 0; i < graph.Size; i++ {
		if graph.InDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var order []int
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, vc := range graph.Edges[v] {
			graph.InDegree[vc]--
			if graph.InDegree[vc] == 0 {
				queue = append(queue, vc)
			}
		}
	}
	return order, len(order) == graph.Size
}

type CaseArg struct {
	N     int
	Names []string
}

func solve(ca CaseArg) string {
	graph := &Graph{
		Size:     26,
		Edges:    make([][]int, 26),
		InDegree: make([]int, 26),
	}

	var prev string
	for _, cur := range ca.Names {
		minL := len(cur)
		if len(prev) < minL {
			minL = len(prev)
		}
		ok := false
		for i := 0; i < minL; i++ {
			if prev[i] != cur[i] {
				from := int(prev[i] - 'a')
				to := int(cur[i] - 'a')
				graph.AddEdge(from, to)
				ok = true
				break
			}
		}
		if !ok && len(prev) > len(cur) {
			return "Impossible"
		}
		prev = cur
	}
	order, acyclic := graph.TopSort()
	if !acyclic {
		return "Impossible"
	}
	var ans strings.Builder
	done := make([]bool, 26)
	for _, c := range order {
		ans.WriteByte(byte('a' + c))
		done[c] = true
	}
	for i, d := range done {
		if !d {
			ans.WriteByte(byte('a' + i))
		}
	}
	return ans.String()
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	for i := 0; i < ca.N; i++ {
		var name string
		fmt.Fscan(in, &name)
		ca.Names = append(ca.Names, name)
	}
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
