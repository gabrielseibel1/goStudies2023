package paths

import "github.com/gabrielseibel1/goStudies2023/graph"

// https://leetcode.com/problems/all-paths-from-source-to-target/

func AllPathsSourceTarget(graph graph.Graph) [][]int {
	if len(graph) == 0 {
		return make([][]int, 0)
	}

	paths := allPaths(graph, []int{0})

	// paths := allPaths(graph)
	return paths
}

// func allPaths(graph graph.Graph) [][]int {
// 	paths := make([][]int, 0)

// 	// start queue with path that has just first node
// 	q := make([][]int, 1)
// 	q[0] = []int{0}

// 	for len(q) > 0 {
// 		// dequeue a path
// 		p := q[0]
// 		q = q[1:]

// 		// get tip of path
// 		tip := p[len(p)-1]

// 		// check if found destination
// 		if tip == len(graph)-1 {
// 			paths = append(paths, p)
// 			continue
// 		}

// 		// enqueue paths with new neighbors
// 		for _, n := range graph[tip].Neighbors {
// 			np := make([]int, len(p))
// 			copy(np, p)
// 			q = append(q, append(np, n))
// 		}
// 	}

// 	return paths
// }

func allPaths(graph graph.Graph, path []int) [][]int {
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
	paths := make([][]int, 0, len(graph[i].Neighbors))
	for _, n := range graph[i].Neighbors {
		nPath := append(path, n)
		nPaths := allPaths(graph, nPath)
		paths = append(paths, nPaths...)
	}
	return paths
}
