// C. Table Decorations
// https://codeforces.com/problemset/problem/478/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Color struct {
	Index int
	Count int
}

func max(arr ...int) int {
	var ret int = math.MinInt
	for _, num := range arr {
		if num > ret {
			ret = num
		}
	}
	return ret
}

func solve(ca CaseArg) int {
	var ans int
	m_color := max(ca.Colors[0], ca.Colors[1], ca.Colors[2])
	s_color := ca.Colors[0] + ca.Colors[1] + ca.Colors[2]
	if m_color > 2*(s_color-m_color) {
		ans = s_color - m_color
	} else {
		ans = s_color / 3
	}
	return ans
}

type CaseArg struct {
	Colors []int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	ca.Colors = make([]int, 3)
	fmt.Fscan(in, &ca.Colors[0], &ca.Colors[1], &ca.Colors[2])
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
