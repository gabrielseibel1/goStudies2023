package bridges

import (
	"math"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

var count int

func FindBridges(g graph.Graph) []*graph.Edge {
	low := make([]int, len(g))
	ids := make([]int, len(g))
	vis := make([]bool, len(g))

	bridges := make([]*graph.Edge, 0)
	dfs(g, vis, ids, low, 0, -1, &bridges)

	return bridges
}

func dfs(g graph.Graph, vis []bool, ids []int, low []int, at, parent int, bridges *[]*graph.Edge) {
	vis[at] = true
	count++
	ids[at] = count
	low[at] = count

	for i, to := range g[at].Neighbors {
		if to == parent {
			continue
		}

		if ok := vis[to]; ok {
			low[at] = int(math.Min(float64(low[at]), float64(ids[to])))
		} else {
			dfs(g, vis, ids, low, to, at, bridges)
			low[at] = int(math.Min(float64(low[at]), float64(low[to])))
			if ids[at] < low[to] {
				*bridges = append(*bridges, &graph.Edge{
					From: at,
					To:   to,
					Dist: g[at].Distances[i],
				})
			}
		}
	}
}
