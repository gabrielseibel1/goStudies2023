package bipartite

import "github.com/gabrielseibel1/goStudies2023/graph"

type color uint8

const (
	blank color = iota
	blue
	red
)

func IsBipartite(g graph.Graph) bool {
	colors := make([]color, len(g))
	for node := range g {
		if colors[node] != blank { // seen
			continue
		}
		if ok := dfs(g, node, nextColor(colors[node]), &colors); !ok {
			return false
		}
	}
	return true
}

func dfs(g graph.Graph, node int, requiredColor color, colors *[]color) bool {
	switch (*colors)[node] {
	case requiredColor:
		return true
	case nextColor(requiredColor):
		return false
	case blank:
		(*colors)[node] = requiredColor
	}

	for _, node := range g[node].Neighbors {
		if ok := dfs(g, node, nextColor(requiredColor), colors); !ok {
			return false
		}
	}
	return true
}

func nextColor(c color) color {
	switch c {
	case blank:
		return red
	case red:
		return blue
	case blue:
		return red
	}
	return blank
}
