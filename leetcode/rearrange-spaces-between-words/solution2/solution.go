package solution1

import "strings"

// problem: https://leetcode.cn/problems/rearrange-spaces-between-words/

func reorderSpaces(text string) (ans string) {
	words := strings.Fields(text)
	spaces := strings.Count(text, " ")
	if len(words) == 1 {
		ans += words[0] + strings.Repeat(" ", spaces)
		return
	}
	midspace := strings.Repeat(" ", spaces/(len(words)-1))
	for i := 0; i < len(words)-1; i++ {
		ans += words[i] + midspace
	}
	ans += words[len(words)-1] + strings.Repeat(" ", spaces%(len(words)-1))
	return
}
