package solution

import (
	"github.com/emirpasic/gods/maps/treemap"
)

// Item ...
type Item struct {
	Index int
	Value int
}

func oddEvenJumps(A []int) int {
	odds := make([]bool, len(A))
	evens := make([]bool, len(A))
	odds[len(A)-1] = true
	evens[len(A)-1] = true
	tm := treemap.NewWithIntComparator()
	tm.Put(A[len(A)-1], len(A)-1)
	count := 1
	for i := len(A) - 2; i >= 0; i-- {
		v := A[i]
		floorKey, floorValue := tm.Floor(v)
		ceilingKey, ceilingValue := tm.Ceiling(v)
		if ceilingKey != nil {
			cv, _ := ceilingValue.(int)
			if evens[cv] == true {
				odds[i] = true
				count++
			}
		}
		if floorKey != nil {
			fv, _ := floorValue.(int)
			if odds[fv] == true {
				evens[i] = true
			}
		}
		tm.Put(v, i)
	}

	return count
}

// Run ...
func Run(A []int) int {
	return oddEvenJumps(A)
}
