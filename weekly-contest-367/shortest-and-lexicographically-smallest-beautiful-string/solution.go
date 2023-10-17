package shortestandlexicographicallysmallestbeautifulstring

func shortestBeautifulSubstring(s string, k int) string {
	onearr := make([]int, len(s)+1)
	one := 0
	i := 1
	for _, c := range s {
		if c == '1' {
			one++
		}
		onearr[i] = one
		i++
	}
	var ans string = s + "1"
	for i := 1; i < len(onearr); i++ {
		for j := i; j < len(onearr); j++ {
			if onearr[j]-onearr[i-1] == k {
				sub := s[i-1 : j]
				if len(sub) < len(ans) || (len(sub) == len(ans) && sub < ans) {
					ans = sub
				}
			}
		}
	}
	if ans == s+"1" {
		return ""
	}
	return ans
}
