package validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	num := len(pushed)
	arr := make([]int, len(pushed))
	arr_idx := -1
	pushed_idx := 0
	for poi := 0; poi < num; poi++ {
		pov := popped[poi]

		if arr_idx != -1 && arr[arr_idx] == pov {
			arr_idx--
			continue
		}

		var isFound bool
		for ; pushed_idx < num; pushed_idx++ {
			puv := pushed[pushed_idx]
			if puv == pov {
				isFound = true
				pushed_idx++
				break
			}
			arr_idx++
			arr[arr_idx] = puv
		}
		if !isFound {
			return false
		}
	}
	return true
}
