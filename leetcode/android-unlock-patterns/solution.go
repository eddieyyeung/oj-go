package androidunlockpatterns

import "math"

var minStep int
var maxStep int

// numberOfPatterns https://leetcode-cn.com/problems/android-unlock-patterns/
func numberOfPatterns(m int, n int) int {
	minStep = m
	maxStep = n
	v := [3][3]bool{}
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bt(&v, 1, i, j, &count)
		}
	}
	return count
}

func bt(v *[3][3]bool, step int, x int, y int, count *int) {
	if step >= minStep && step <= maxStep {
		*count++
	}
	if step >= maxStep {
		return
	}
	v[x][y] = true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if v[i][j] {
				continue
			}
			diffX := int(math.Abs(float64(i) - float64(x)))
			diffY := int(math.Abs(float64(j) - float64(y)))
			if diffX < 2 && diffY < 2 {
				if !v[i][j] {
					bt(v, step+1, i, j, count)
				}
				continue
			}
			sumX := i + x
			sumY := j + y
			midX := sumX / 2
			midY := sumY / 2
			if sumX%2 == 0 && sumY%2 == 0 {
				if v[midX][midY] {
					bt(v, step+1, i, j, count)
				}
				continue
			}
			bt(v, step+1, i, j, count)
		}
	}
	v[x][y] = false
}
