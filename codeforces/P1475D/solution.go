// D. Cleaning the Phone
// https://codeforces.com/problemset/problem/1475/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type CaseArg struct {
	N        int
	M        int
	Memories []int
	Points   []int
}

func solve(ca CaseArg) int {
	indexes := make([]int, ca.N)
	var total_points int = 0
	for i := 0; i < len(indexes); i++ {
		indexes[i] = i
		total_points += ca.Points[i]
	}
	sort.Slice(indexes, func(i, j int) bool {
		if ca.Points[indexes[i]] < ca.Points[indexes[j]] {
			return true
		} else if ca.Points[indexes[i]] > ca.Points[indexes[j]] {
			return false
		} else {
			if ca.Points[indexes[i]] == 1 {
				return ca.Memories[indexes[i]] < ca.Memories[indexes[j]]
			} else {
				return ca.Memories[indexes[i]] > ca.Memories[indexes[j]]
			}
		}
	})
	var j int = 0
	var free_m int = 0
	var cost_p int = 0
	var ans int = math.MaxInt
	for i := 0; i < len(indexes); i++ {
		var m int = ca.Memories[indexes[i]]
		var p int = ca.Points[indexes[i]]
		free_m += m
		cost_p += p
		for free_m >= ca.M {
			ans = min(ans, cost_p)
			free_m -= ca.Memories[indexes[j]]
			cost_p -= ca.Points[indexes[j]]
			j++
		}
	}
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func runCase(in io.Reader, out io.Writer) {
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N, &ca.M)
		ca.Memories = make([]int, ca.N)
		ca.Points = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.Memories[j])
		}
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.Points[j])
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
