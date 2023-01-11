package topological

func IsTopologicallySorted(graph [][]int, sorting []int) bool {
	seen := make(map[int]struct{})

	for _, curr := range sorting {
		seen[curr] = struct{}{}
		for _, n := range graph[curr] {
			if _, s := seen[n]; s {
				return false
			}
		}
	}
	return true
}

func TopSort(graph [][]int) []int {
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

func traverseStacking(graph [][]int, node int, seen map[int]struct{}) []int {
	if _, s := seen[node]; s {
		return make([]int, 0)
	}

	seen[node] = struct{}{}
	order := make([]int, 0)
	for _, n := range graph[node] {
		order = append(traverseStacking(graph, n, seen), order...)
	}
	return append([]int{node}, order...)
}
