package backwardsreadprimes

// BackwardsPrime https://www.codewars.com/kata/5539fecef69c483c5a000015/train/go
func BackwardsPrime(start int, stop int) []int {
	m := map[int]bool{}
	i := start
	result := []int{}
	if i < 11 {
		i = 11
	}
	for ; i <= stop; i++ {
		if m[i] == true {
			result = append(result, i)
			continue
		}
		if isPrime(i) == true {
			reverse := reverseInt(i)
			if reverse != i && isPrime(reverse) == true {
				m[i] = true
				m[reverse] = true
				result = append(result, i)
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func reverseInt(num int) int {
	n := 0
	for num > 0 {
		r := num % 10
		n *= 10
		n += r
		num /= 10
	}
	return n
}
