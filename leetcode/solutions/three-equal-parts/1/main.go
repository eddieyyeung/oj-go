package solution

import "fmt"

func getReal(a []int, startIndex int) (int, []int) {
	i := startIndex
	for ; i < len(a)-1; i++ {
		if a[i] == 1 {
			break
		}
	}
	return i, a[i:]
}

// ThreeEqualParts https://leetcode-cn.com/problems/three-equal-parts/
func threeEqualParts(A []int) []int {
	oneNum := 0
	for _, v := range A {
		oneNum += v
	}
	if oneNum%3 != 0 {
		return []int{-1, -1}
	}
	if oneNum == 0 {
		return []int{0, 2}
	}
	partion := oneNum / 3
	partionA := 0
	zeroNum := 0

	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 1 {
			break
		}
		if A[i] == 0 {
			zeroNum++
		}
	}
	for i := 1; i <= len(A); i++ {
		partionA += A[i-1]
		if partionA > partion {
			break
		}
		if partionA != partion {
			continue
		}
		for z := 0; z < zeroNum; z++ {
			if A[i+z] != 0 {
				return []int{-1, -1}
			}
		}
		i += zeroNum
		a := A[:i]
		_, realA := getReal(a, 0)
		partionB := 0
		for j := i + 1; j <= len(A); j++ {
			if partionB > partion {
				break
			}
			partionB += A[j-1]
			if partionB != partion {
				continue
			}
			for z := 1; z <= zeroNum; z++ {
				if A[j+z-1] != 0 {
					return []int{-1, -1}
				}
			}
			j += zeroNum
			b := A[i:j]
			_, realB := getReal(b, 0)
			c := A[j:]
			_, realC := getReal(c, 0)
			if len(realA) != len(realB) || len(realA) != len(realC) {
				continue
			}
			for u, v := range realA {
				if v != realB[u] || v != realC[u] {
					return []int{-1, -1}
				}
			}
			return []int{i - 1, j}
		}
	}
	return []int{-1, -1}
}

// Run ...
func Run(A []int) []int {
	return threeEqualParts(A)
}

func init() {
	arr := []int{0, 0, 0, 0, 0}
	fmt.Println(threeEqualParts(arr))
}
