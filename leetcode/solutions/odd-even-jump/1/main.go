package solution

import (
	"container/list"
)

type Item struct {
	index  int
	isOdd  bool
	isEven bool
}
type VisitEven struct {
	isVisit bool
	isReach bool
}

type VisitOdd struct {
	isVisit bool
	isReach bool
}

func oddEvenJumps(A []int) int {
	visitedEven := make([]bool, len(A))
	visitedOdd := make([]bool, len(A))
	visitedEven[len(A)-1] = true
	visitedOdd[len(A)-1] = true
	queue := list.New()
	queue.PushBack(Item{len(A) - 1, true, true})
	count := 1
	for queue.Len() != 0 {
		e := queue.Front()
		item, _ := e.Value.(Item)
		a := A[item.index]
		min := -1
		max := 100001
		for i := item.index + 1; i < len(A); i++ {
			b := A[i]
			if b > a && b < max {
				max = b
			}
			if b < a && b > min {
				min = b
			}
		}
		for i := item.index - 1; i >= 0; i-- {
			b := A[i]
			ve := visitedEven[i]
			vo := visitedOdd[i]
			if ve == false && item.isOdd == true && b > min && b <= a {
				count++
				queue.PushBack(Item{i, false, true})
				visitedEven[i] = true
			}
			if vo == false && item.isEven == true && b >= a && b < max {
				queue.PushBack(Item{i, true, false})
				visitedOdd[i] = true
			}

			if b <= a && b > min {
				min = b
			}
			if b >= a && b < max {
				max = b
			}
		}
		queue.Remove(e)
	}
	return count
}

func Run(A []int) int {
	return oddEvenJumps(A)
}
