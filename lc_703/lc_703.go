package lc703

/*
Design a class to find the kth largest element in a stream.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.

int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.
*/

import (
	"container/heap"
)

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{nums, k}
	heap.Init(&h)
	h.prune()
	return h
}

func (h *KthLargest) Add(val int) int {
	heap.Push(h, val)
	h.prune()
	return h.nums[0]
}

func (h *KthLargest) prune() {
	for h.Len() > h.k {
		heap.Pop(h)
	}
}

// Len is the number of elements in the collection.
func (h *KthLargest) Len() int {
	return len(h.nums)
}

// Less reports whether the element with index i must sort before the element with index j.
func (h *KthLargest) Less(i int, j int) bool {
	return h.nums[i] < h.nums[j]
}

// Swap swaps the elements with indexes i and j.
func (h *KthLargest) Swap(i int, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *KthLargest) Push(x any) {
	h.nums = append(h.nums, x.(int))
}

func (h *KthLargest) Pop() any {
	if len(h.nums) == 0 {
		panic("empty heap")
	}
	defer func() { h.nums = h.nums[0 : len(h.nums)-1] }()
	return h.nums[len(h.nums)-1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
