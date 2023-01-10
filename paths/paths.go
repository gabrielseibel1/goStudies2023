package paths

func AllPathsSourceTarget(graph [][]int) [][]int {
	if len(graph) == 0 {
		return make([][]int, 0)
	}

	paths := allPaths(graph, []int{0})
	return paths
}

func allPaths(graph [][]int, path []int) [][]int {
	// pop from path after done, on return
	defer func() {
		path = path[:len(path)-1]
	}()

	// current node
	i := path[len(path)-1]

	// reached end
	if i == len(graph)-1 {
		p := make([]int, len(path))
		copy(p, path)
		return [][]int{p}
	}

	// travel to neighbors
	paths := make([][]int, 0, len(graph[i]))
	for _, n := range graph[i] {
		nPath := append(path, n)
		nPaths := allPaths(graph, nPath)
		paths = append(paths, nPaths...)
	}
	return paths
}
