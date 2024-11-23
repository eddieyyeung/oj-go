package solution

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/description/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	type Node struct {
		Val  int
		Next *Node
	}
	m := make([]*Node, n)
	root := &Node{}
	p := root
	dist := n - 1
	for i := 0; i < n; i++ {
		node := Node{
			Val:  i,
			Next: nil,
		}
		p.Next = &node
		p = p.Next
		m[i] = &node
	}
	var ans []int = make([]int, 0, len(queries))
	for _, query := range queries {
		x, y := query[0], query[1]
		if m[x] == nil {
			ans = append(ans, dist)
			continue
		}
		xnext := m[x].Next
		for xnext.Val < y {
			m[x].Next = xnext.Next
			m[xnext.Val] = nil
			xnext.Next = nil
			xnext = m[x].Next
			dist--
		}
		ans = append(ans, dist)
	}
	return ans
}
