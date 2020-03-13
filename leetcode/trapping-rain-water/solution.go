// Package trapping_rain_water ...
// https://leetcode-cn.com/problems/trapping-rain-water/
package trapping_rain_water

import "math"

type node struct {
	Val    int
	Weight int
	Pre    *node
	Next   *node
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	p := &node{
		height[0],
		1,
		nil,
		nil,
	}
	p.Next = &node{
		height[1],
		1,
		p,
		nil,
	}
	p = p.Next
	rain := 0
	for i := 2; i < len(height); i++ {
		n := &node{
			height[i],
			1,
			nil,
			nil,
		}
		w := 0
		for p.Pre != nil {
			if p.Pre.Val >= p.Val && p.Val <= n.Val {
				w += p.Weight
				rain += w * (min(p.Pre.Val, n.Val) - p.Val)
				p = p.Pre
				p.Next = nil
			} else {
				break
			}
		}
		n.Pre = p
		n.Weight += w
		p.Next = n
		p = n
	}
	return rain
}

func min(a ...int) int {
	min := math.MaxInt64
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
