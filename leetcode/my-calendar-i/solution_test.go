package mycalendari

import (
	"fmt"
	"testing"
)

func TestMyCalendar_Book(t *testing.T) {
	mc := Constructor()
	fmt.Println(mc.Book(37, 50))
	fmt.Println(mc.Book(33, 50))
	fmt.Println(mc.Book(4, 17))
	fmt.Println(mc.Book(35, 48))
	fmt.Println(mc.Book(8, 25))
}
