package topological

import "github.com/gabrielseibel1/goStudies2023/graph"

func IsTopologicallySorted(g graph.Graph, sorting []int) bool {
	seen := make([]bool, len(g))

	for _, curr := range sorting {
		seen[curr] = true
		for _, n := range g[curr].Neighbors {
			if seen[n] {
				return false
			}
		}
	}

	for i := range g {
		if !seen[i] {
			return false
		}
	}
	return true
}

func TopSort(g graph.Graph) []int {
	seen := make([]bool, len(g))

	reverse := make([]int, 0, len(g))
	for i := range g {
		if !seen[i] {
			dfs(g, i, seen, &reverse)
		}
	}

	order := make([]int, len(reverse))
	for i, v := range reverse {
		order[len(order)-1-i] = v
	}
	return order
}

func dfs(g graph.Graph, i int, seen []bool, reverse *[]int) {
	seen[i] = true

	for _, n := range g[i].Neighbors {
		if seen[n] {
			continue
		}
		dfs(g, n, seen, reverse)
	}

	*reverse = append(*reverse, i)
}
