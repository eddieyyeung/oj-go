package solution

var (
	lena int
	numa int
	lenb int
	numb int
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) == 0 && len(nums2) == 1 {
		return float64(nums2[0])
	}
	if len(nums1) == 1 && len(nums2) == 0 {
		return float64(nums1[0])
	}
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		lena = totalLen/2 - 1
		lenb = totalLen / 2
	} else {
		lena = totalLen / 2
		lenb = totalLen / 2
	}
	fn1(nums1, nums2)
	//fmt.Println("-------")
	//fmt.Println("numa", numa)
	//fmt.Println("numb", numb)
	return float64(numa+numb) / 2
}

func fn1(nums1 []int, nums2 []int) {
	//fmt.Println("------")
	//fmt.Println("lena", lena)
	//fmt.Println("lenb", lenb)
	//fmt.Println("nums1", nums1)
	//fmt.Println("nums2", nums2)
	//fmt.Println("numa", numa)
	//fmt.Println("numb", numb)
	if len(nums1) == 0 && len(nums2) > 0 {
		if lena >= 0 {
			numa = nums2[lena]
		}
		if lenb >= 0 {
			numb = nums2[lenb]
		}
		return
	}
	if len(nums2) == 0 && len(nums1) > 0 {
		if lena >= 0 {
			numa = nums1[lena]
		}
		if lenb >= 0 {
			numb = nums1[lenb]
		}
		return
	}
	num1 := nums1[0]
	if num1 < nums2[0] {
		if lena >= 0 {
			lena = lena - 1
			if lena <= -1 {
				numa = num1
			}
		}
		if lenb >= 0 {
			lenb = lenb - 1
			if lenb <= -1 {
				numb = num1
			}
		}

		if lena <= -1 && lenb <= -1 {
			return
		}
		nums1 = nums1[1:]
		fn1(nums2, nums1)
	} else {
		idx := fn2(0, len(nums2)-1, nums2, num1)
		//fmt.Println("idx", idx)
		if idx >= lena && lena >= 0 {
			numa = nums2[lena]
			lena = -1
		}
		if idx >= lenb && lenb >= 0 {
			numb = nums2[lenb]
			lenb = -1
		}
		if lena <= -1 && lenb <= -1 {
			return
		}
		if lena >= 0 {
			lena = lena - idx - 1
		}
		if lenb >= 0 {
			lenb = lenb - idx - 1
		}
		nums2 = nums2[idx+1:]
		fn1(nums2, nums1)
	}
}

func fn2(left int, right int, nums []int, target int) int {
	if left <= right {
		mid := (left + right) / 2
		//fmt.Println("mid", mid)
		if nums[mid] <= target {
			left = mid + 1
			return fn2(left, right, nums, target)
		} else {
			right = mid - 1
			return fn2(left, right, nums, target)
		}
	}
	return (left + right) / 2
}
