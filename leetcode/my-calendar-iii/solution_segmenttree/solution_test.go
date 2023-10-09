package mycalendariii

import (
	"fmt"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	mc := Constructor()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(50, 60))
	fmt.Println(mc.Book(10, 40))
	fmt.Println(mc.Book(5, 15))
	fmt.Println(mc.Book(5, 10))
	fmt.Println(mc.Book(25, 55))
}
