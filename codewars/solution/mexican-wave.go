package solution

import "unicode"

// MaxicanWave https://www.codewars.com/kata/mexican-wave/go
func MaxicanWave(words string) []string {
	wave := []string{}
	for i, v := range words {
		if v == ' ' {
			continue
		}
		wave = append(wave, words[:i]+string(unicode.ToUpper(v))+words[i+1:])
	}
	return wave
}
