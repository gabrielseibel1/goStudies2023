package graph

import "fmt"

type GNode struct {
	Distances []int
	Neighbors []int
}

type Graph []GNode

var SquareDAG = []GNode{
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

var TrapezoidDAG = []GNode{
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

var LongDAG = []GNode{
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
