package solution

import (
	"container/list"
	"sort"
)

type Item struct {
	index int
	value int
}
type SortBy []Item

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].value < a[j].value }

func runWithMonotonousStack(items []Item) []int {
	s := list.New()
	arr := make([]int, len(items))
	for _, item := range items {
		for s.Len() != 0 {
			back := s.Back()
			v, _ := back.Value.(Item)
			if v.index <= item.index {
				arr[v.index] = item.index
				s.Remove(back)
			} else {
				break
			}
		}
		s.PushBack(item)
	}
	return arr
}

func oddEvenJumps(A []int) int {
	items := make([]Item, len(A))
	for i := 0; i < len(A); i++ {
		items[i] = Item{
			i,
			A[i],
		}
	}
	sort.SliceStable(items, func(i int, j int) bool {
		return items[i].value < items[j].value
	})
	oddNext := runWithMonotonousStack(items)
	sort.SliceStable(items, func(i int, j int) bool {
		return items[i].value > items[j].value
	})
	evenNext := runWithMonotonousStack(items)
	odds := make([]bool, len(A))
	evens := make([]bool, len(A))
	odds[len(A)-1] = true
	evens[len(A)-1] = true
	count := 1
	for i := len(A) - 2; i >= 0; i-- {
		if oddNext[i] != 0 && evens[oddNext[i]] == true {
			odds[i] = true
			count++
		}
		if evenNext[i] != 0 && odds[evenNext[i]] == true {
			evens[i] = true
		}
	}
	return count
}

func Run(A []int) int {
	return oddEvenJumps(A)
}
