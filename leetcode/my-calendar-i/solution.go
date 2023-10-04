package mycalendari

// 执行用时 84 ms 击败 67.74%
// 消耗内存 7.7 MB 击败 6.45%
import (
	"sort"
)

type MyCalendar struct {
	Books [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	idx := sort.Search(len(mc.Books), func(i int) bool {
		return start < mc.Books[i][0]
	})
	if idx > 0 {
		if mc.Books[idx-1][1] > start {
			return false
		}
	}
	if idx < len(mc.Books) {
		if mc.Books[idx][0] < end {
			return false
		}
	}
	mc.Books = append(mc.Books[:idx], append([][]int{{start, end}}, mc.Books[idx:]...)...)
	return true
}
