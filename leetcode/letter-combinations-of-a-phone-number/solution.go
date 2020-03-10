// Package lettercombinationsofaphonenumber
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
package lettercombinationsofaphonenumber

var digitMap map[string][]string = map[string][]string{
	"1": []string{},
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	re := []string{}
	if len(digits) == 0 {
		return re
	}
	bt(digits, "", &re)
	return re
}

func bt(digits string, lc string, re *[]string) {
	if len(digits) == 0 {
		*re = append(*re, lc)
		return
	}
	letters := digitMap[string(digits[0])]
	for _, l := range letters {
		bt(digits[1:], lc+l, re)
	}
}
