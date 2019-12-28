package solution

import (
	"strconv"
)

// RotateForAMax https://www.codewars.com/kata/rotate-for-a-max/go
func RotateForAMax(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	len := len(s)
	max := n
	for i := 0; i < len-1; i++ {
		s = s[:i] + s[i+1:] + s[i:i+1]
		k, _ := strconv.ParseInt(s, 10, 64)
		if k > max {
			max = k
		}
	}
	return max
}
