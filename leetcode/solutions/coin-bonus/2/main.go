package solution

import (
	"container/list"
	"fmt"
)

var totalCoins [50000]uint64

type Member struct {
	Leader         *Member
	SubFollowers   *list.List
	Coins          uint64
	FollowersNum   int
	FollowersCoins uint64
	Level          int
}

func addCoinsWithFollowers(m *Member, coin uint64) {
	m.Coins += coin
	m.FollowersCoins += uint64(m.FollowersNum) * coin
	for e := m.SubFollowers.Front(); e != nil; e = e.Next() {
		n, _ := e.Value.(*Member)
		addCoinsWithFollowers(n, coin)
	}
}

func setFollowersNum(m *Member) {
	m.FollowersNum = 0
	for e := m.SubFollowers.Front(); e != nil; e = e.Next() {
		n, _ := e.Value.(*Member)
		setFollowersNum(n)
		m.FollowersNum += n.FollowersNum + 1
	}
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	var result []int
	var members = make([]Member, n)
	for i := 0; i < n; i++ {
		members[i].SubFollowers = list.New()
		members[i].Level = i
	}
	for i := 0; i < len(leadership); i++ {
		members[leadership[i][0]-1].SubFollowers.PushFront(&members[leadership[i][1]-1])
		members[leadership[i][1]-1].Leader = &members[leadership[i][0]-1]
	}
	setFollowersNum(&members[0])
	// print(members)

	for i := 0; i < len(operations); i++ {
		o1 := operations[i][0]
		o2 := operations[i][1]
		if o1 == 1 {
			coin := uint64(operations[i][2])
			p := &members[o2-1]
			p.Coins += coin
			for p.Leader != nil {
				p = p.Leader
				p.FollowersCoins += coin
			}
		} else if o1 == 2 {
			p := &members[o2-1]
			coin := uint64(operations[i][2])
			addedFollowersCoins := coin*uint64(p.FollowersNum) + coin
			addCoinsWithFollowers(p, coin)
			for p.Leader != nil {
				p = p.Leader
				p.FollowersCoins += addedFollowersCoins
			}
		} else {
			m := members[o2-1]
			totalCoins := (m.Coins + m.FollowersCoins) % 1000000007
			result = append(result, int(totalCoins))
		}
	}
	// for i := 0; i < len(members); i++ {
	// 	fmt.Println(i+1, ": ", members[i].Coins, "|", members[i].FollowersCoins)
	// }
	return result
}

func print(members []Member) {
	for i := 0; i < len(members); i++ {
		m := members[i]
		fmt.Print(i+1, "|", m.FollowersNum, ": ")
		for e := m.SubFollowers.Front(); e != nil; e = e.Next() {
			n, _ := e.Value.(*Member)
			fmt.Print(n.Level+1, "|", n.FollowersNum, " ")
		}
		fmt.Println()
	}
}

func Run(n int, leadership [][]int, operations [][]int) []int {
	return bonus(n, leadership, operations)
}
