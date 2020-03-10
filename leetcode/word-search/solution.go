// Package wordsearch ...
// https://leetcode-cn.com/problems/word-search/
package wordsearch

type move struct {
	x int
	y int
}

type pos struct {
	x int
	y int
}

var moves []move = []move{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var visited [][]bool

func exist(board [][]byte, word string) bool {
	visited = make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := range board {
		for j := range board[i] {
			r := bt(0, word, board, pos{i, j})
			if r == true {
				return true
			}
		}
	}
	return false
}

func bt(idx int, word string, board [][]byte, p pos) bool {
	if idx >= len(word) || board[p.x][p.y] != word[idx] {
		return false
	}
	if len(word) == idx+1 {
		return true
	}
	visited[p.x][p.y] = true
	for _, move := range moves {
		newP := pos{
			p.x + move.x,
			p.y + move.y,
		}
		if newP.x < 0 || newP.y < 0 || newP.x >= len(board) || newP.y >= len(board[newP.x]) {
			continue
		}
		if visited[newP.x][newP.y] == true {
			continue
		}
		// visited[newP.x][newP.y] = true
		if bt(idx+1, word, board, newP) == true {
			return true
		}
		// visited[newP.x][newP.y] = false
	}
	visited[p.x][p.y] = false
	return false
}
