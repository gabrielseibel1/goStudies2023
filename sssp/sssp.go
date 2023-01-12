package sssp

// single source shortest path

import (
	"math"

	"github.com/gabrielseibel1/goStudies2023/graph"
	"github.com/gabrielseibel1/goStudies2023/topological"
)

func ShortestPath(graph graph.Graph) []int {
	distances := make([]int, len(graph))
	distances[0] = 0
	for i := 1; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}

	ordering := topological.TopSort(graph)

	for _, node := range ordering {
		if distances[node] == math.MaxInt {
			continue
		}
		for i, neig := range graph[node].Neighbors {
			d := distances[neig]
			dd := distances[node] + graph[node].Distances[i]
			distances[neig] = int(math.Min(float64(d), float64(dd)))
		}
	}
	return distances
}
