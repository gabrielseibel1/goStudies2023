package lc_84_largest_rectangle

import (
	"math"
)

func largestRectangleArea(heights []int) int {
	l, h := largestRectDims(heights, 0)
	return l * h
}

func largestRectDims(heights []int, oh int) (ll int, lh int) {
	if len(heights) == 0 {
		return 0, oh
	}
	if len(heights) == 1 {
		return len(heights), heights[0] + oh
	}
	l, h := 0, 0
	mini := subtractMin(heights)
	parts := partitions(heights)
	for i := range parts {
		d1, d2 := largestRectDims(heights[parts[i][0]:parts[i][1]+1], mini+oh)
		if d1*d2 > l*h {
			l, h = d1, d2
		}
	}
	if len(heights)*(mini+oh) >= l*h {
		return len(heights), mini + oh
	} else {
		return l, h
	}
}

func partitions(heights []int) [][]int {
	var parts [][]int
	var s int
	for i, h := range heights {
		if h > 0 && i > 0 && heights[i-1] == 0 {
			s = i
		}
		if h == 0 && i > 0 && heights[i-1] > 0 {
			parts = append(parts, []int{s, i - 1})
		} else if h > 0 && i == len(heights)-1 {
			parts = append(parts, []int{s, i})
		}
	}
	return parts
}

func subtractMin(heights []int) int {
	mini := math.MaxInt
	for _, h := range heights {
		if h < mini {
			mini = h
		}
	}
	for i := range heights {
		heights[i] -= mini
	}
	return mini
}
