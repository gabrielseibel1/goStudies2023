package quicksort

import "math"

func QuickSort(s []int) []int {
	s = append(s, math.MaxInt)
	quickSort(&s, 0, len(s)-1)
	return s[:len(s)-1]
}

func quickSort(s *[]int, from, to int) {
	if from == to {
		return
	}

	pivotIndex := partition(s, from, to)
	quickSort(s, 0, pivotIndex)
	quickSort(s, pivotIndex+1, to)
}

func partition(s *[]int, i, j int) int {
	pivot := (*s)[(i+j)/2]
	for {
		for ; (*s)[i] < pivot; i++ {
		}
		for ; (*s)[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		swap := (*s)[i]
		(*s)[i] = (*s)[j]
		(*s)[j] = swap
	}
	return j
}
