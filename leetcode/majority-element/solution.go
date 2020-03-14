package majority_element

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
