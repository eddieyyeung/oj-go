package crawler_log_folder

import "strings"

func minOperations(logs []string) (ans int) {
	for _, log := range logs {
		if strings.HasPrefix(log, "./") {
			continue
		}
		if strings.HasPrefix(log, "../") {
			if ans == 0 {
				continue
			}
			ans--
			continue
		}
		ans++
	}

	return ans
}
