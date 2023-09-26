package binaryindexedtree

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	na := Constructor([]int{4, 7, 2, 3, 6, 4, 2, 3, 4, 6, 7, 8, 8, 2, 3})
	fmt.Println(na)
}
