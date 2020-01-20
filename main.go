package main

import (
	"fmt"
	solution "oj-go/leetcode/solutions/unique-paths-ii/1"
)

func main() {
	fmt.Println(solution.Run([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}))
	fmt.Println(solution.Run([][]int{
		[]int{1, 0},
	}))
	fmt.Println(solution.Run([][]int{
		[]int{0, 1},
	}))
}
