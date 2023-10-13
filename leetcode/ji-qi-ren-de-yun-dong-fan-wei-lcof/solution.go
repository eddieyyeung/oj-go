package jiqirendeyundongfanweilcof

func digit(n int) int {
	var ans int
	for n != 0 {
		ans += n % 10
		n = n / 10
	}
	return ans
}

func wardrobeFinishing(m int, n int, cnt int) int {
	var ans int = 1
	v := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		v[i] = make([]int, n+1)
	}
	v[1][1] = 1

	for i := 0; i < m; i++ {
		di := digit(i)
		for j := 0; j < n; j++ {
			if v[i+1][j+1] != 0 {
				continue
			}
			if v[i][j+1] == 1 || v[i+1][j] == 1 {
				dj := digit(j)
				if di+dj <= cnt {
					ans++
					v[i+1][j+1] = 1
				} else {
					v[i+1][j+1] = 2
				}
			} else {
				v[i+1][j+1] = 2
			}
		}
	}
	return ans
}
