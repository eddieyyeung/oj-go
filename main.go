package main

import (
	"fmt"
	solution "oj-go/leetcode/solutions/longest-palindromic-substring/1"
)

func main() {
	fmt.Println(solution.Run("babad"))
	fmt.Println(solution.Run("cbbd"))
	fmt.Println(solution.Run("a"))
	fmt.Println(solution.Run("bb"))
	fmt.Println(solution.Run(""))
	fmt.Println(solution.Run("abcdasdfghjkldcba"))
	fmt.Println(solution.Run("222020221"))
	fmt.Println(solution.Run("SQQSYYSQQS"))
	// fmt.Println(solution.Run("babaddtattarrattatddetartrateedredividerb"))
}
