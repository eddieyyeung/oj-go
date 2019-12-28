package solution

// MakeTheDeadfishSwim https://www.codewars.com/kata/make-the-deadfish-swim/go
func MakeTheDeadfishSwim(data string) []int {
	result := []int{}
	n := 0
	for _, c := range data {
		switch c {
		case 'i':
			n++
		case 'd':
			n--
		case 's':
			n *= n
		case 'o':
			result = append(result, n)
		}
	}
	return result
}
