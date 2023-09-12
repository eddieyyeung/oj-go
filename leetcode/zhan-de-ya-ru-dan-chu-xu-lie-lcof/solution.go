package zhandeyarudanchuxulielcof

func validateStackSequences(pushed []int, popped []int) bool {
	var q []int
	i := 0
	for _, pop := range popped {
		for len(q) <= 0 || q[len(q)-1] != pop {
			if i >= len(pushed) {
				return false
			}
			q = append(q, pushed[i])
			i++
		}
		q = q[:len(q)-1]
	}
	return true
}
