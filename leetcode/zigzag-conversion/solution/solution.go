package solution

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	size := len(s)
	b := make([]byte, 0, len(s))
	div := numRows - 1
	tg := div * 2
	for i := 0; i < numRows; i++ {
		remain := i % div
		gap := tg - (remain * 2)
		idx := i
		for idx < size {
			b = append(b, s[idx])
			idx = idx + gap
			if remain != 0 {
				gap = tg - gap
			}
		}
	}
	return string(b)
}
