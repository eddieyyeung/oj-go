package shortestpathwithalternatingcolors

type State struct {
	RedShortestPath  int
	BlueShortestPath int
}

var (
	redHash  [][]int
	blueHash [][]int
	states   []State
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redHash = make([][]int, n)
	blueHash = make([][]int, n)

	for _, edge := range redEdges {
		redHash[edge[0]] = append(redHash[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		blueHash[edge[0]] = append(blueHash[edge[0]], edge[1])
	}
	states = make([]State, n)
	for i := 1; i < n; i++ {
		states[i].BlueShortestPath = -1
		states[i].RedShortestPath = -1
	}

	for i := 0; i < n; i++ {
		update(i)
	}

	output := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			output[i] = 0
			continue
		}
		min := -1
		if states[i].BlueShortestPath != -1 {
			min = states[i].BlueShortestPath
		}
		if states[i].RedShortestPath != -1 {
			if min == -1 {
				min = states[i].RedShortestPath
			} else {
				if states[i].RedShortestPath < min {
					min = states[i].RedShortestPath
				}
			}
		}
		output[i] = min
	}
	return output
}

func update(n int) {
	redDistNodes := redHash[n]
	blueDistNodes := blueHash[n]
	curRedShortestPath := states[n].RedShortestPath
	curBlueShortestPath := states[n].BlueShortestPath
	if curRedShortestPath != -1 {
		for _, node := range redDistNodes {
			path := curRedShortestPath + 1
			if states[node].BlueShortestPath == -1 {
				states[node].BlueShortestPath = path
				update(node)
			} else {
				if path < states[node].BlueShortestPath {
					states[node].BlueShortestPath = path
					update(node)
				}
			}
		}
	}
	if curBlueShortestPath != -1 {
		for _, node := range blueDistNodes {
			path := curBlueShortestPath + 1
			if states[node].RedShortestPath == -1 {
				states[node].RedShortestPath = path
				update(node)
			} else {
				if path < states[node].RedShortestPath {
					states[node].RedShortestPath = path
					update(node)
				}
			}
		}
	}
}
