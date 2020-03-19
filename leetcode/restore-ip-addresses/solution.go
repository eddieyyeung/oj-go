// Package restore_ip_addresses ...
// https://leetcode-cn.com/problems/restore-ip-addresses/
package restore_ip_addresses

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	m := map[string]bool{}
	ss := []string{}
	bt(s, "", 1, 0, m, &ss)
	return ss
}

func bt(s string, ip string, idx int, pos int, m map[string]bool, ss *[]string) {
	if idx == 4 {
		numstr := s[pos:]
		num, _ := strconv.Atoi(numstr)
		if strconv.Itoa(num) != numstr {
			return
		}
		if num >= 0 && num <= 255 {
			if _, ok := m[ip+numstr]; !ok {
				*ss = append(*ss, ip+numstr)
			}
			m[ip+numstr] = true
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if i+pos >= len(s) {
			return
		}
		numstr := s[pos : i+pos]
		num, _ := strconv.Atoi(numstr)
		if strconv.Itoa(num) != numstr {
			continue
		}
		if num >= 0 && num <= 255 {
			bt(s, ip+numstr+".", idx+1, pos+i, m, ss)
		}
	}
}
