package lc295

import "container/heap"

type MedianFinder struct {
	lowerHalfMax *maxIntHeap
	upperHalfMin *minIntHeap
	even         bool
}

func Constructor() MedianFinder {
	f := MedianFinder{
		lowerHalfMax: new(maxIntHeap),
		upperHalfMin: new(minIntHeap),
		even:         true,
	}
	heap.Init(f.lowerHalfMax)
	heap.Init(f.upperHalfMin)
	return f
}

func (f *MedianFinder) AddNum(num int) {
	if f.even {
		heap.Push(f.lowerHalfMax, num)
		heap.Push(f.upperHalfMin, heap.Pop(f.lowerHalfMax))
	} else {
		heap.Push(f.upperHalfMin, num)
		heap.Push(f.lowerHalfMax, heap.Pop(f.upperHalfMin))
	}
	f.even = !f.even
}

func (f *MedianFinder) FindMedian() float64 {
	if f.even {
		a, b := heap.Pop(f.lowerHalfMax).(int), heap.Pop(f.upperHalfMin).(int)
		defer heap.Push(f.lowerHalfMax, a)
		defer heap.Push(f.upperHalfMin, b)
		return float64(a+b) / 2
	} else {
		a := heap.Pop(f.upperHalfMin).(int)
		defer heap.Push(f.upperHalfMin, a)
		return float64(a)
	}
}

type minIntHeap []int

type maxIntHeap []int

func (m *minIntHeap) Len() int {
	return len(*m)
}

func (m *minIntHeap) Less(i int, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minIntHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minIntHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minIntHeap) Pop() any {
	defer func() {
		*m = (*m)[:m.Len()-1]
	}()
	return (*m)[m.Len()-1]
}

func (m *maxIntHeap) Len() int {
	return len(*m)
}

func (m *maxIntHeap) Less(i int, j int) bool {
	return (*m)[j] < (*m)[i]
}

func (m *maxIntHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxIntHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxIntHeap) Pop() any {
	defer func() {
		*m = (*m)[:m.Len()-1]
	}()
	return (*m)[m.Len()-1]
}
