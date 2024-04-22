// D. Carousel
// https://codeforces.com/problemset/problem/1328/D
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

type Item struct {
	Index int
	Value int
}

type MaxHeap []Item

// Len implements heap.Interface.
func (m MaxHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MaxHeap) Less(i int, j int) bool {
	return m[i].Value > m[j].Value
}

// Pop implements heap.Interface.
func (m *MaxHeap) Pop() any {
	old := *m
	l := len(old)
	v := old[l-1]
	*m = old[:l-1]
	return v
}

// Push implements heap.Interface.
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(Item))
}

// Swap implements heap.Interface.
func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = &MaxHeap{}

func solve(ca CaseArg) string {
	var ans string
	var pre int = -1
	var win_turn int = -1
	var h MaxHeap
	heap.Init(&h)
	for i := 0; i < ca.N; i++ {
		heap.Push(&h, Item{
			Index: i,
			Value: ca.As[i],
		})
	}
	for {
		if len(h) == 0 {
			break
		}
		last := heap.Pop(&h).(Item)
		if last.Index == pre {
			if len(h) == 0 {
				break
			}
			last2 := heap.Pop(&h).(Item)
			pre = last2.Index
			win_turn = -win_turn
			last2.Value--
			if last2.Value != 0 {
				heap.Push(&h, last2)
			}
			heap.Push(&h, last)
		} else {
			pre = last.Index
			win_turn = -win_turn
			last.Value--
			if last.Value != 0 {
				heap.Push(&h, last)
			}
		}
	}
	if win_turn == 1 {
		ans = "T"
	} else {
		ans = "HL"
	}

	return ans
}

type CaseArg struct {
	N  int
	As []int
}

func runCase(in io.Reader, out io.Writer) {
	var case_num int
	fmt.Fscan(in, &case_num)
	for i := 0; i < case_num; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.As = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.As[j])
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
