package topological

import "github.com/gabrielseibel1/goStudies2023/graph"

func IsTopologicallySorted(graph graph.Graph, sorting []int) bool {
	seen := make(map[int]struct{})

	for _, curr := range sorting {
		seen[curr] = struct{}{}
		for _, n := range graph[curr].Neighbors {
			if _, s := seen[n]; s {
				return false
			}
		}
	}
	return true
}

func TopSort(graph graph.Graph) []int {
	seen := make(map[int]struct{})

	revOrder := make([]int, 0, len(graph))
	for node := range graph {
		revOrder = append(traverseStacking(graph, node, seen), revOrder...)
	}

	order := make([]int, 0, len(graph))
	for _, v := range revOrder {
		order = append(order, v)
	}
	return order
}

func traverseStacking(graph graph.Graph, node int, seen map[int]struct{}) []int {
	if _, s := seen[node]; s {
		return make([]int, 0)
	}

	seen[node] = struct{}{}
	order := make([]int, 0)
	for _, n := range graph[node].Neighbors {
		order = append(traverseStacking(graph, n, seen), order...)
	}
	return append([]int{node}, order...)
}
