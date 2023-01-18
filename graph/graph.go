package graph

import "fmt"

type Edge struct {
	From int
	To   int
	Dist int
}

func (e *Edge) String() string {
	return fmt.Sprintf("{ %d -(%d)-> %d } ", e.From, e.Dist, e.To)
}

type GNode struct {
	Distances []int
	Neighbors []int
}

type Graph []GNode

var SquareDG = []GNode{
	{
		Distances: []int{1, 2},
		Neighbors: []int{1, 2},
	},
	{
		Distances: []int{1},
		Neighbors: []int{3},
	},
	{
		Distances: []int{-1},
		Neighbors: []int{3},
	},
	{
		Distances: []int{},
		Neighbors: []int{},
	},
}

var Square = []GNode{
	{
		Distances: []int{1, 2},
		Neighbors: []int{1, 2},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{3, 0},
	},
	{
		Distances: []int{-1, 2},
		Neighbors: []int{3, 0},
	},
	{
		Distances: []int{1, -1},
		Neighbors: []int{1, 2},
	},
}

var TrapezoidDG = []GNode{
	{
		Distances: []int{4, 1, 1},
		Neighbors: []int{4, 3, 1},
	},
	{
		Distances: []int{-1, 2},
		Neighbors: []int{3, 2},
	},
	{
		Distances: []int{-2},
		Neighbors: []int{3},
	},
	{
		Distances: []int{2},
		Neighbors: []int{4},
	},
	{
		Distances: []int{},
		Neighbors: []int{},
	},
}

var Trapezoid = []GNode{
	{
		Distances: []int{4, 1, 1},
		Neighbors: []int{4, 3, 1},
	},
	{
		Distances: []int{-1, 2, 1},
		Neighbors: []int{3, 2, 0},
	},
	{
		Distances: []int{-2, 2},
		Neighbors: []int{3, 1},
	},
	{
		Distances: []int{2, 1, -1, -2},
		Neighbors: []int{4, 0, 1, 2},
	},
	{
		Distances: []int{4, 2},
		Neighbors: []int{0, 3},
	},
}

var NoBridge = []GNode{
	{
		Distances: []int{1, 1},
		Neighbors: []int{1, 2},
	},
	{
		Distances: []int{1, 1, 1},
		Neighbors: []int{2, 3, 0},
	},
	{
		Distances: []int{1, 1, 1},
		Neighbors: []int{3, 1, 0},
	},
	{
		Distances: []int{1, 1, 1, 1},
		Neighbors: []int{4, 5, 1, 2},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{5, 3},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{3, 4},
	},
}

var Bridge = []GNode{
	{
		Distances: []int{1, 1},
		Neighbors: []int{1, 2},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{2, 0},
	},
	{
		Distances: []int{1, 1, 1},
		Neighbors: []int{3, 1, 0},
	},
	{
		Distances: []int{1, 1, 1, 1},
		Neighbors: []int{4, 6, 7, 2},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{5, 3},
	},
	{
		Distances: []int{1},
		Neighbors: []int{4},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{3, 7},
	},
	{
		Distances: []int{1, 1},
		Neighbors: []int{3, 6},
	},
}

var LongDG = []GNode{
	{
		Distances: []int{2, 4},
		Neighbors: []int{1, 2},
	},
	{
		Distances: []int{5},
		Neighbors: []int{3},
	},
	{
		Distances: []int{-3},
		Neighbors: []int{3},
	},
	{
		Distances: []int{3, 4},
		Neighbors: []int{6, 12},
	},
	{
		Distances: []int{-1, 2, 6},
		Neighbors: []int{1, 3, 5},
	},
	{
		Distances: []int{4, 11},
		Neighbors: []int{7, 8},
	},
	{
		Distances: []int{-4, 5},
		Neighbors: []int{7, 11},
	},
	{
		Distances: []int{9, 1},
		Neighbors: []int{9, 10},
	},
	{
		Distances: []int{2},
		Neighbors: []int{7},
	},
	{
		Distances: []int{},
		Neighbors: []int{},
	},
	{
		Distances: []int{},
		Neighbors: []int{},
	},
	{
		Distances: []int{-2},
		Neighbors: []int{10},
	},
	{
		Distances: []int{11},
		Neighbors: []int{11},
	},
}

func (gn GNode) String() string {
	b := []byte{}
	for i, n := range gn.Neighbors {
		b = fmt.Appendf(b, "\n\t-(%d)-> [%d]\n", gn.Distances[i], n)
	}
	return string(b)
}

func (g Graph) String() string {
	b := []byte{}
	for i, n := range g {
		b = fmt.Appendf(b, "[%d]:%s", i, n)
	}
	return string(b)
}
