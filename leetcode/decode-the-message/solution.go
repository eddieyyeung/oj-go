package decodethemessage

func decodeMessage(key string, message string) string {
	cipher := make([]int, 26)
	i := 0
	for _, r := range key {
		if r != ' ' {
			if cipher[r-'a'] == 0 {
				i++
				cipher[r-'a'] = i
			}
		}
	}
	output := make([]byte, 0, len(message))
	for _, r := range message {
		if r != ' ' {
			output = append(output, byte(cipher[r-'a']+'a'-1))
		} else {
			output = append(output, byte(r))
		}
	}
	return string(output)
}
