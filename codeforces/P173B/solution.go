// B. Chamber of Secrets
// https://codeforces.com/problemset/problem/173/B
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
)

type Node struct {
	X         int16
	Y         int16
	Direction int16
	Step      int16
}

var fx []int16 = []int16{0, 1, 0, -1}
var fy []int16 = []int16{1, 0, -1, 0}

type Deque struct {
	L []any
	R []any
}

func (d *Deque) Empty() bool {
	return len(d.L) == 0 && len(d.R) == 0
}

func (d *Deque) PushL(n any) {
	if cap(d.L) == len(d.L) {
		s := make([]any, len(d.L), len(d.L)+1_000_000)
		copy(s, d.L)
		d.L = s
		runtime.GC()
	}
	d.L = append(d.L, n)
}

func (d *Deque) PushR(n any) {
	if cap(d.R) == len(d.R) {
		s := make([]any, len(d.R), len(d.R)+1_000_000)
		copy(s, d.R)
		d.R = s
		runtime.GC()
	}
	d.R = append(d.R, n)
}

func (d *Deque) PopL() (n any) {
	if len(d.L) > 0 {
		n, d.L = d.L[len(d.L)-1], d.L[:len(d.L)-1]
	} else {
		n, d.R = d.R[0], d.R[1:]
	}
	return
}

func solve(ca CaseArg) int16 {
	visited := [1000][1000][4]int16{}
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			for k := 0; k < 4; k++ {
				visited[i][j][k] = math.MaxInt16
			}
		}
	}
	d := Deque{}
	d.PushL(Node{
		X:         0,
		Y:         0,
		Direction: 0,
		Step:      0,
	})
	visited[0][0][0] = 1
	for !d.Empty() {
		n := d.PopL().(Node)
		if n.X == ca.N-1 && n.Y == ca.M-1 && n.Direction == 0 {
			return n.Step
		}
		nx := n.X + fx[n.Direction]
		ny := n.Y + fy[n.Direction]
		if nx >= 0 && nx < ca.N && ny >= 0 && ny < ca.M {
			if n.Step < visited[nx][ny][n.Direction] {
				visited[nx][ny][n.Direction] = n.Step
				d.PushL(Node{
					X:         nx,
					Y:         ny,
					Direction: n.Direction,
					Step:      n.Step,
				})
			}
		}
		if ca.Grid[n.X][n.Y] == '#' {
			for i := int16(0); i < 4; i++ {
				if i == n.Direction {
					continue
				}
				if n.Step+1 < visited[n.X][n.Y][i] {
					visited[n.X][n.Y][i] = n.Step + 1
					d.PushR(Node{
						X:         n.X,
						Y:         n.Y,
						Direction: i,
						Step:      n.Step + 1,
					})
				}
			}
		}
	}

	return -1
}

type CaseArg struct {
	N, M int16
	Grid [][]byte
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.N)
	fmt.Fscan(in, &ca.M)
	ca.Grid = make([][]byte, ca.N)
	for i := int16(0); i < ca.N; i++ {
		fmt.Fscan(in, &ca.Grid[i])
	}
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
