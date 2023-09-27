package queuereconstructionbyheight

import (
	"container/list"
	"fmt"
	"sort"
)

type Item struct {
	Idx int
	H   int
	K   int
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	var ans [][]int
	l := list.New()
	for i, p := range people {
		l.PushBack(&Item{
			Idx: i,
			H:   p[0],
			K:   p[1],
		})
	}
	for l.Len() > 0 {
		for e := l.Front(); e != nil; e = e.Next() {
			item := e.Value.(*Item)
			if item.K == 0 {
				ans = append(ans, people[item.Idx])
				for j := l.Front(); j != e; j = j.Next() {
					item2 := j.Value.(*Item)
					item2.K--
				}
				l.Remove(e)
				break
			}
		}
	}
	return ans
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		item := e.Value.(*Item)
		fmt.Println(item)
	}
	fmt.Println("----")
}
