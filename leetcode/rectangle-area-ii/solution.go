package rectangle_area_ii

import (
	"sort"
)

// problem: https://leetcode.cn/problems/rectangle-area-ii/

type Item struct {
	YCoord int
	XLeft  int
	XRight int
	Idx    int
	IsLow  bool
}

func rectangleArea(rectangles [][]int) int {
	var items []Item
	for i, rectangle := range rectangles {
		items = append(items, Item{
			YCoord: rectangle[1],
			XLeft:  rectangle[0],
			XRight: rectangle[2],
			Idx:    i,
			IsLow:  true,
		}, Item{
			YCoord: rectangle[3],
			XLeft:  rectangle[0],
			XRight: rectangle[2],
			Idx:    i,
			IsLow:  false,
		})

	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].YCoord < items[j].YCoord {
			return true
		} else if items[i].YCoord > items[j].YCoord {
			return false
		} else {
			return items[i].XLeft < items[j].XLeft
		}
	})
	var arr []Item
	var sum int
	ycoord := 0
	for _, item := range items {
		moved := item.YCoord - ycoord
		var tw int
		left := 0
		rm := -1
		for i, item2 := range arr {
			if !item.IsLow && item.Idx == item2.Idx {
				rm = i
			}
			if item2.XRight <= left {
				continue
			}
			if item2.XLeft >= left {
				tw += item2.XRight - item2.XLeft
				left = item2.XRight
				continue
			}
			tw += item2.XRight - left
			left = item2.XRight
		}
		sum += moved * tw
		if !item.IsLow {
			for i := rm; i < len(arr)-1; i++ {
				arr[i] = arr[i+1]
			}
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr, item)
			sort.Slice(arr, func(i, j int) bool {
				return arr[i].XLeft < arr[j].XLeft
			})
		}
		ycoord = item.YCoord
	}
	return int(sum % (1e9 + 7))
}
