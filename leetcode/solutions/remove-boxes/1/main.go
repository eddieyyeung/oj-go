package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func remove(boxes []int, set map[string]int) int {
	if len(boxes) == 1 {
		return 1
	}
	result := 0
	for i, j := 0, 0; i < len(boxes); i++ {
		for j < len(boxes) && boxes[i] == boxes[j] {
			j++
		}
		arr := make([]int, len(boxes)-(j-i))
		a := boxes[0:i]
		b := boxes[j:]
		k := 0
		for i := 0; i < len(a); i++ {
			arr[k] = a[i]
			k++
		}
		for i := 0; i < len(b); i++ {
			arr[k] = b[i]
			k++
		}
		sArr := make([]string, len(arr))
		for i, v := range arr {
			sArr[i] = strconv.Itoa(v)
		}
		s := strings.Join(sArr, ",")
		val, exists := set[s]
		var add int
		if exists == true {
			fmt.Println("val", val)
			add = val
		} else {
			add = remove(arr, set)
			set[s] = add
		}
		r := (j-i)*(j-i) + add
		if r > result {
			result = r
		}
		i = j - 1
	}
	return result
}

func removeBoxes(boxes []int) int {
	set := make(map[string]int)
	return remove(boxes, set)
}

// Run ...
func Run(boxes []int) int {
	return removeBoxes(boxes)
}
