package pickpeaks

// PosPeaks ...
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
func PickPeaks(array []int) PosPeaks {
	pospeaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
	pre := 0
	cur := 1
	for i := 2; i < len(array); i++ {
		if array[i] != array[cur] {
			if array[pre] < array[cur] && array[cur] > array[i] {
				pospeaks.Pos = append(pospeaks.Pos, cur)
				pospeaks.Peaks = append(pospeaks.Peaks, array[cur])
			}
			pre = cur
			cur = i
		}
	}
	return pospeaks
}
