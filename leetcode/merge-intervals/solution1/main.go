package solution1

import (
	"math"
	"sort"
)

type Intervals [][]int

// Len implements sort.Interface.
func (ints *Intervals) Len() int {
	return len([][]int(*ints))
}

// Less implements sort.Interface.
func (ints *Intervals) Less(i int, j int) bool {
	return [][]int(*ints)[i][0] < [][]int(*ints)[j][0]
}

// Swap implements sort.Interface.
func (ints *Intervals) Swap(i int, j int) {
	[][]int(*ints)[i], [][]int(*ints)[j] = [][]int(*ints)[j], [][]int(*ints)[i]
}

var _ sort.Interface = &Intervals{}

func MaxInt(nums ...int) int {
	i := math.MinInt
	for _, num := range nums {
		if num > i {
			i = num
		}
	}
	return i
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	ints := Intervals(intervals)
	sort.Sort(&ints)

	var result [][]int
	prev := ints[0]

	for i := 1; i < len(ints); i++ {
		cur := ints[i]
		if cur[0] > prev[1] {
			result = append(result, []int{prev[0], prev[1]})
			prev[0] = cur[0]
			prev[1] = cur[1]
		} else {
			prev[1] = MaxInt(prev[1], cur[1])
		}
	}
	result = append(result, []int{prev[0], prev[1]})
	return result
}
