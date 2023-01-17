package prim

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func MinimumSpanningTree(g graph.Graph) (dist int, edges []*Edge) {
	seen := make(map[int]struct{})
	edges = make([]*Edge, 0, len(g)-1)
	pq := make(priorityQueue, 0)

	// init node 0
	edges = append(edges, &Edge{
		from: -1,
		to:   0,
		dist: 0,
	})
	seen[0] = struct{}{}
	pushEdges(&pq, seen, g, 0)

	// do prims algorithm
	for pq.Len() > 0 && len(edges) < len(g) {
		var e *Edge = heap.Pop(&pq).(*Edge)
		dist += e.dist

		if _, ok := seen[e.to]; ok {
			continue
		}

		edges = append(edges, e)

		seen[e.to] = struct{}{}

		pushEdges(&pq, seen, g, e.to)
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
		heap.Push(pq, &Edge{
			from: i,
			to:   node.Neighbors[j],
			dist: node.Distances[j],
		})
	}
}

type Edge struct {
	from int
	to   int
	dist int
}

func (e *Edge) String() string {
	return fmt.Sprintf("{ %d -(%d)-> %d } ", e.from, e.dist, e.to)
}

type priorityQueue []*Edge

func (pq *priorityQueue) Push(e any) {
	*pq = append(*pq, e.(*Edge))
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
	return (*pq)[i].dist < (*pq)[j].dist
}

func (pq *priorityQueue) Swap(i, j int) {
	s := (*pq)[i]
	(*pq)[i] = (*pq)[j]
	(*pq)[j] = s
}
