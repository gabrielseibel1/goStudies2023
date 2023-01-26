package mergesort

// https://leetcode.com/problems/sort-list

import (
	"math"

	"github.com/gabrielseibel1/goStudies2023/list"
)

func MergeSort(head *list.Node) *list.Node {
	if head == nil {
		return head
	}
	// traverse list to count elements once
	n := countElements(head)
	return mergeSort(head, n)
}

func countElements(l *list.Node) int {
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}

func mergeSort(h *list.Node, n int) *list.Node {
	if n == 1 {
		h.Next = nil
		return h
	}

	// divide the list in two and recurse
	r := h
	for i := 0; i < n/2; i++ {
		r = r.Next
	}
	sl := mergeSort(h, n/2)
	sr := mergeSort(r, int(math.Ceil(float64(n)/2)))

	// merge the two lists
	var ln *list.Node
	if sl.Val <= sr.Val {
		ln = sl
		sl = sl.Next
	} else {
		ln = sr
		sr = sr.Next
	}
	mergedSortedHead := ln
	for !(sl == nil && sr == nil) { // only done when two are nil
		// one of them could be nil though
		slv := math.MaxInt
		if sl != nil {
			slv = sl.Val
		}
		srv := math.MaxInt
		if sr != nil {
			srv = sr.Val
		}

		if slv <= srv {
			ln.Next = sl
			sl = sl.Next
		} else {
			ln.Next = sr
			sr = sr.Next
		}
		ln = ln.Next
	}
	ln.Next = nil
	return mergedSortedHead
}
