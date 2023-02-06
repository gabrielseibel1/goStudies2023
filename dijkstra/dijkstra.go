package dijkstra

// solve https://leetcode.com/problems/network-delay-time

import (
	"container/heap"
	"math"
)

func NetworkDelayTime(times [][]int, n int, k int) int {
	dists := dijkstra(times, n, k)

	max := maxOf(dists)
	if max == math.MaxInt {
		return -1
	}
	return max
}

func maxOf(s []int) int {
	max := math.MinInt
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func adjListFromEdgeList(edges [][]int, n int) [][][]int {
	g := make([][][]int, n+1)
	for _, e := range edges {
		from, to, dist := e[0], e[1], e[2]
		g[from] = append(g[from], []int{to, dist})
	}
	return g
}

func dijkstra(times [][]int, n int, k int) []int {
	g := adjListFromEdgeList(times, n)

	dists := make([]int, n+1)
	for i := range dists {
		dists[i] = math.MaxInt
	}
	dists[k] = 0

	pq := make(edgePQ, 0)
	pq.push(&edge{
		node: k,
		dist: 0,
	})

	for pq.Len() > 0 {
		e := pq.pop()

		if e.dist > dists[e.node] {
			continue
		}

		for _, adj := range g[e.node] {
			v, d := adj[0], adj[1]
			sourceToVUsingE := dists[e.node] + d
			if sourceToVUsingE < dists[v] {
				dists[v] = sourceToVUsingE
				pq.Push(&edge{
					node: v,
					dist: sourceToVUsingE,
				})
			}
		}
	}

	return dists[1:]
}

type edge struct {
	node int
	dist int
}

type edgePQ []*edge

func (pq *edgePQ) push(e *edge) {
	heap.Push(pq, e)
}

func (pq *edgePQ) pop() *edge {
	return heap.Pop(pq).(*edge)
}

func (pq edgePQ) Len() int {
	return len(pq)
}

func (pq edgePQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq *edgePQ) Swap(i, j int) {
	tmp := (*pq)[i]
	(*pq)[i] = (*pq)[j]
	(*pq)[j] = tmp
}

func (pq *edgePQ) Push(x interface{}) {
	e := x.(*edge)
	*pq = append(*pq, e)
}

func (pq *edgePQ) Pop() interface{} {
	n := len(*pq) - 1
	e := (*pq)[n]
	(*pq)[n] = nil
	*pq = (*pq)[:n]
	return e
}
