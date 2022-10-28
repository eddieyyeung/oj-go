package remove_one_element_to_make_the_array_strictly_increasing

// 2 1 2 3
func canBeIncreasing(nums []int) bool {
	// fmt.Println("ltor", ltor(nums))
	// fmt.Println("rtol", rtol(nums))
	return ltor(nums) || rtol(nums)
}

func ltor(nums []int) bool {
	var arr []int
	var c int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		al := len(arr)
		arr = append(arr, num)
		var isChange bool
		for j := al; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isChange = true
			} else if arr[j] == arr[j-1] {
				isChange = true
				break
			} else {
				break
			}
		}
		if isChange {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return true
}

func rtol(nums []int) bool {
	var arr []int
	var c int
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		al := len(arr)
		arr = append(arr, num)
		var isChange bool
		for j := al; j > 0; j-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				isChange = true
			} else if arr[j] == arr[j-1] {
				isChange = true
				break
			} else {
				break
			}
		}
		if isChange {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return true
}
