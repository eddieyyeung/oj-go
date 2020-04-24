// Package permutation_sequence ...
// https://leetcode-cn.com/problems/permutation-sequence/
package permutation_sequence

import (
	"container/list"
	"strconv"
)

func getPermutation(n int, k int) string {
	l := list.New()
	var p []int
	result := ""
	p = append(p, 1)
	for i := 1; i <= n; i++ {
		l.PushBack(i)
		p = append(p, p[i-1]*i)
	}
	for i := n - 1; i >= 1; i-- {
		pn := p[i]
		pos := (k - 1) / pn
		e := l.Front()
		for t := 0; t < pos; t++ {
			e = e.Next()
		}
		v, _ := e.Value.(int)
		result += strconv.Itoa(v)
		l.Remove(e)
		k = k - pos*pn
	}
	v, _ := l.Front().Value.(int)
	result += strconv.Itoa(v)
	return result
}
