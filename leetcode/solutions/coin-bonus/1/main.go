package solution

import (
	"container/list"
	"fmt"
)

var totalCoins [50000]uint64

func addCoins(ms []*list.List, coins []uint64, member int, coin uint64) {
	coins[member] += coin
	for e := ms[member].Front(); e != nil; e = e.Next() {
		n, _ := e.Value.(int)
		addCoins(ms, coins, n, coin)
	}
}
func getCoins(ms []*list.List, coins []uint64, member int) uint64 {
	totalCoins := coins[member]
	for e := ms[member].Front(); e != nil; e = e.Next() {
		n, _ := e.Value.(int)
		totalCoins += getCoins(ms, coins, n)
	}

	return totalCoins
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	var result []int
	var members = make([]*list.List, n)
	for i := 0; i < n; i++ {
		members[i] = list.New()
	}
	for i := 0; i < len(leadership); i++ {
		members[leadership[i][0]-1].PushFront(leadership[i][1] - 1)
	}
	coins := make([]uint64, n)
	for i := 0; i < len(operations); i++ {
		o1 := operations[i][0]
		o2 := operations[i][1]
		if o1 == 1 {
			coins[o2-1] += uint64(operations[i][2])
		} else if o1 == 2 {
			addCoins(members, coins, o2-1, uint64(operations[i][2]))
		} else {
			totalCoins := getCoins(members, coins, o2-1) % 1000000007
			result = append(result, int(totalCoins))
		}
	}
	return result
}

func print(ms []*list.List) {
	for i := 0; i < len(ms); i++ {
		fmt.Print(i+1, ":")
		for e := ms[i].Front(); e != nil; e = e.Next() {
			n, _ := e.Value.(int)
			fmt.Printf("%d ", n+1)
		}
		fmt.Println()
	}
}

func Run(n int, leadership [][]int, operations [][]int) []int {
	return bonus(n, leadership, operations)
}
