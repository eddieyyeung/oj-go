// D. Good Substrings
// https://codeforces.com/problemset/problem/271/D
package main

import (
	"bufio"
	"fmt"
	"index/suffixarray"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// solution3 use suffixarray
func solve(ca CaseArg) int {
	var n int = len(ca.S)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(ca.S))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && ca.S[i+h] == ca.S[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	sum := make([]int, n+1)
	for i, b := range ca.S {
		sum[i+1] = sum[i] + int(ca.Letters[b-'a']&15^1)
	}
	var ans int
	for i, p := range sa {
		p := int(p)
		for j := p + height[i] + 1; j <= n; j++ {
			if sum[j]-sum[p] <= ca.K {
				ans++
			}
		}
	}
	return ans
}

// solution2
// use trie tree
// type Node struct {
// 	Children [26]*Node
// }

// func solve(ca CaseArg) int {
// 	root := &Node{}
// 	var ans int

// 	for i := 0; i < len(ca.S); i++ {
// 		cnt, p := 0, root
// 		for j := i; j < len(ca.S); j++ {
// 			idx := ca.S[j] - 'a'
// 			if ca.Letters[idx] == '0' {
// 				cnt++
// 			}
// 			if cnt > ca.K {
// 				break
// 			}
// 			if p.Children[idx] == nil {
// 				p.Children[idx] = &Node{}
// 				ans++
// 			}
// 			p = p.Children[idx]
// 		}
// 	}
// 	return ans
// }

// solution 1
// use cache to store found string
// func solve(ca CaseArg) int {
// 	fits := make(map[string]bool)
// 	cnt_notfit := 0
// 	var t int = 0

// 	setFits := func(l, r int) {
// 		for i := l; i <= r; i++ {
// 			is_continue := true
// 			for j := r + 1; j >= i+1 && is_continue; j-- {
// 				if fits[ca.S[i:j]] {
// 					is_continue = false
// 				}
// 				fits[ca.S[i:j]] = true
// 			}
// 		}
// 	}

// 	for i := 0; i < len(ca.S); i++ {
// 		if ca.Letters[ca.S[i]-'a'] == '0' {
// 			cnt_notfit++
// 			for cnt_notfit > ca.K {
// 				if ca.Letters[ca.S[t]-'a'] == '0' {
// 					cnt_notfit--
// 				}
// 				t++
// 			}
// 		}
// 		setFits(t, i)
// 	}
// 	return len(fits)
// }

type CaseArg struct {
	S       string
	Letters string
	K       int
}

func runCase(in io.Reader, out io.Writer) {
	var ca CaseArg
	fmt.Fscan(in, &ca.S)
	fmt.Fscan(in, &ca.Letters)
	fmt.Fscan(in, &ca.K)
	fmt.Fprintln(out, solve(ca))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runCase(in, out)
}
