// 752. Open the Lock
// https://leetcode.cn/problems/open-the-lock/description/
package solution

// 0000 -> 1000 -> .... -> 8888
func openLock(deadends []string, target string) int {
	type Item struct {
		Value string
		Step  int
	}
	getNexts := func(str string) []string {
		b := []byte(str)
		var rst []string
		for i, c := range b {
			b[i] = c - 1
			if b[i] < '0' {
				b[i] = '9'
			}
			rst = append(rst, string(b))
			b[i] = c + 1
			if b[i] > '9' {
				b[i] = '0'
			}
			rst = append(rst, string(b))
			b[i] = c
		}
		return rst
	}

	const start = "0000"
	if target == start {
		return 0
	}
	dead := make(map[string]bool)
	seen := map[string]bool{start: true}
	for _, d := range deadends {
		dead[d] = true
	}
	if dead[start] {
		return -1
	}
	q := []Item{{start, 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, next := range getNexts(curr.Value) {
			if !seen[next] && !dead[next] {
				if next == target {
					return curr.Step + 1
				}
				seen[next] = true
				q = append(q, Item{next, curr.Step + 1})
			}
		}
	}
	return -1
}
