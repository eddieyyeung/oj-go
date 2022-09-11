package solution1

// problem: https://leetcode.cn/problems/rearrange-spaces-between-words/

func reorderSpaces(text string) string {
	strs := []string{}
	s := 0
	isNextWord := false
	totalSpace := 0
	for i, c := range text {
		if c == ' ' {
			totalSpace++
			if isNextWord {
				isNextWord = false
				strs = append(strs, text[s:i])
			}
		} else {
			if !isNextWord {
				isNextWord = true
				s = i
			}
		}
	}
	if isNextWord {
		strs = append(strs, text[s:])
	}
	ans := ""
	if len(strs) == 1 {
		spaces := ""
		for i := 0; i < totalSpace; i++ {
			spaces += " "
		}
		ans += strs[len(strs)-1] + spaces
	} else {
		d := totalSpace / (len(strs) - 1)
		spaces := ""
		for i := 0; i < d; i++ {
			spaces += " "
		}
		r := totalSpace % (len(strs) - 1)
		rspaces := ""
		for i := 0; i < r; i++ {
			rspaces += " "
		}
		for i := 0; i < len(strs)-1; i++ {
			ans += strs[i] + spaces
		}
		ans += strs[len(strs)-1] + rspaces
	}
	return ans
}
