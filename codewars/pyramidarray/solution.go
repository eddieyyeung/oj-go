package pyramidarray

// Pyramid https://www.codewars.com/kata/515f51d438015969f7000013/go
func Pyramid(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			result[i][j] = 1
		}
	}
	return result
}
