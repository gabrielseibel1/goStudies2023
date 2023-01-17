package main

import (
	"fmt"

	"github.com/gabrielseibel1/goStudies2023/graph"
	"github.com/gabrielseibel1/goStudies2023/prim"
)

func main() {
	dist, edges := prim.MinimumSpanningTree(graph.Square)
	fmt.Printf("Dist = %d\n", dist)
	fmt.Println(edges)

	dist, edges = prim.MinimumSpanningTree(graph.Trapezoid)
	fmt.Printf("Dist = %d\n", dist)
	fmt.Println(edges)
}
