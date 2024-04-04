// C. Hamburgers
// https://codeforces.com/contest/371/problem/C
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(ca CaseArg) int64 {
	cnts := [3]int64{}
	for _, c := range ca.Letters {
		if c == 'B' {
			cnts[0]++
		} else if c == 'S' {
			cnts[1]++
		} else {
			cnts[2]++
		}
	}
	var ans int64
	for ca.Haves[0] != 0 || ca.Haves[1] != 0 || ca.Haves[2] != 0 {
		ok := true
		prev := ca.Haves
		for i := 0; i < len(ca.Haves); i++ {
			if ca.Haves[i] < cnts[i] {
				if ca.R >= (cnts[i]-ca.Haves[i])*ca.Prices[i] {
					ca.R -= (cnts[i] - ca.Haves[i]) * ca.Prices[i]
					ca.Haves[i] = 0
				} else {
					ok = false
					break
				}
			} else {
				ca.Haves[i] -= cnts[i]
			}
		}
		if ok {
			ans++
			if prev == ca.Haves {
				break
			}
		} else {
			break
		}
	}
	oneprice := cnts[0]*ca.Prices[0] + cnts[1]*ca.Prices[1] + cnts[2]*ca.Prices[2]
	ans += ca.R / oneprice
	return ans
}

type CaseArg struct {
	Letters []byte
	Haves   [3]int64
	Prices  [3]int64
	R       int64
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.Letters)
	fmt.Fscan(in, &ca.Haves[0], &ca.Haves[1], &ca.Haves[2])
	fmt.Fscan(in, &ca.Prices[0], &ca.Prices[1], &ca.Prices[2])
	fmt.Fscan(in, &ca.R)
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
