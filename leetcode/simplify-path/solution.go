// Package simplify_path ...
// https://leetcode-cn.com/problems/simplify-path/
package simplify_path

import (
	"container/list"
	"regexp"
	"strings"
)

func simplifyPath(path string) string {
	dirs := regexp.MustCompile(`/+`).Split(strings.Trim(path, "/"), -1)
	l := list.New()
	for _, dir := range dirs {
		if dir == "." {
			continue
		}
		if dir == ".." {
			if l.Len() == 0 {
				continue
			}
			l.Remove(l.Back())
			continue
		}
		l.PushBack(dir)
	}
	var r strings.Builder
	for l.Len() != 0 {
		e := l.Front()
		dir, _ := e.Value.(string)
		r.WriteString("/")
		r.WriteString(dir)
		l.Remove(e)
	}
	if r.Len() == 0 {
		r.WriteString("/")
	}
	return r.String()
}
