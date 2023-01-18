package prim

import (
	"container/heap"
	"math"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func MinimumSpanningTree(g graph.Graph) (dist int, edges []*graph.Edge) {
	seen := make(map[int]struct{})
	edges = make([]*graph.Edge, 0, len(g)-1)
	pq := make(priorityQueue, 0)

	// init node 0
	edges = append(edges, &graph.Edge{
		From: -1,
		To:   0,
		Dist: 0,
	})
	seen[0] = struct{}{}
	pushEdges(&pq, seen, g, 0)

	// do prims algorithm
	for pq.Len() > 0 && len(edges) < len(g) {
		var e *graph.Edge = heap.Pop(&pq).(*graph.Edge)
		dist += e.Dist

		if _, ok := seen[e.To]; ok {
			continue
		}

		edges = append(edges, e)

		seen[e.To] = struct{}{}

		pushEdges(&pq, seen, g, e.To)
	}

	if len(edges) < len(g) {
		return math.MaxInt, nil
	}

	return
}

func pushEdges(pq *priorityQueue, seen map[int]struct{}, g graph.Graph, i int) {
	node := g[i]
	for j := 0; j < len(node.Distances); j++ {
		if _, ok := seen[node.Neighbors[j]]; ok {
			continue
		}
		heap.Push(pq, &graph.Edge{
			From: i,
			To:   node.Neighbors[j],
			Dist: node.Distances[j],
		})
	}
}

type priorityQueue []*graph.Edge

func (pq *priorityQueue) Push(e any) {
	*pq = append(*pq, e.(*graph.Edge))
}

func (pq *priorityQueue) Pop() any {
	defer func() {
		(*pq)[pq.Len()-1] = nil // avoid memory leak
		*pq = (*pq)[:pq.Len()-1]
	}()
	return (*pq)[pq.Len()-1]
}

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].Dist < (*pq)[j].Dist
}

func (pq *priorityQueue) Swap(i, j int) {
	s := (*pq)[i]
	(*pq)[i] = (*pq)[j]
	(*pq)[j] = s
}
