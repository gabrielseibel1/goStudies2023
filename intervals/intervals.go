package intervals

// https://leetcode.com/problems/non-overlapping-intervals

import "sort"

type sortableIntervals [][]int

func (mh sortableIntervals) Len() int {
	return len(mh)
}

func (mh sortableIntervals) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}

func (mh *sortableIntervals) Swap(i, j int) {
	swap := (*mh)[i]
	(*mh)[i] = (*mh)[j]
	(*mh)[j] = swap
}

func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	var sortedIntervals sortableIntervals = intervals
	sort.Sort(&sortedIntervals)

	count := 0
	prev := 0
	for i := 1; i < len(sortedIntervals); i++ {
		if sortedIntervals[i][0] < sortedIntervals[prev][1] {
			count++
		} else {
			prev = i
		}
	}
	return count
}
