// D. Carousel
// https://codeforces.com/problemset/problem/1328/D
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) (int, []int) {
	ans := make([]int, ca.N)
	ans[0] = 1
	var cnt int
	var a, b, c int
	var is_more_than_two bool
	a = ca.Ts[0]
	var pre int = 0
	for i := 1; i < ca.N; i++ {
		if ca.Ts[i] == ca.Ts[pre] {
			is_more_than_two = true
		}
		pre++
		if ca.Ts[i] != a {
			if b == 0 {
				b = ca.Ts[i]
			} else if b != ca.Ts[i] {
				c = ca.Ts[i]
			}
		}
	}

	if b == 0 {
		for i := 0; i < ca.N; i++ {
			ans[i] = 1
		}
		cnt = 1
	} else if ca.N%2 == 0 {
		pre = 0
		for i := 1; i < ca.N; i++ {
			if ans[pre] == 1 {
				ans[i] = 2
			} else {
				ans[i] = 1
			}
			pre++
		}
		cnt = 2
	} else {
		if is_more_than_two {
			is_dup := false
			pre = -1
			for i := 1; i < ca.N; i++ {
				pre++
				if ca.Ts[pre] == ca.Ts[i] && !is_dup {
					ans[i] = ans[pre]
					is_dup = true
					continue
				}
				if ans[pre] == 1 {
					ans[i] = 2
				} else {
					ans[i] = 1
				}
			}
			cnt = 2
		} else if c == 0 {
			pre = 0
			for i := 1; i < ca.N; i++ {
				if ans[pre] == 1 {
					ans[i] = 2
				} else {
					ans[i] = 1
				}
				pre++
			}
			cnt = 2
		} else {
			pre = 0
			for i := 1; i < ca.N; i++ {
				if ans[pre] == 1 {
					ans[i] = 2
				} else {
					ans[i] = 1
				}
				pre++
			}
			if ca.Ts[ca.N-1] != ca.Ts[0] && ans[ca.N-1] == ans[0] {
				ans[ca.N-1] = 3
				cnt = 3
			} else {
				cnt = 2
			}
		}
	}
	return cnt, ans
}

type CaseArg struct {
	N  int
	Ts []int
}

func runCase(in io.Reader, out io.Writer) {
	var case_num int
	fmt.Fscan(in, &case_num)
	for i := 0; i < case_num; i++ {
		var ca CaseArg
		fmt.Fscan(in, &ca.N)
		ca.Ts = make([]int, ca.N)
		for j := 0; j < ca.N; j++ {
			fmt.Fscan(in, &ca.Ts[j])
		}
		cnt, ans := solve(ca)
		fmt.Fprintln(out, cnt)
		for _, n := range ans {
			fmt.Fprintf(out, "%d ", n)
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
