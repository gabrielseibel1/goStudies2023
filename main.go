package main

import (
	"fmt"

	"github.com/gabrielseibel1/goStudies2023/bridges"
	"github.com/gabrielseibel1/goStudies2023/graph"
)

func main() {
	// b1 := bridges.FindBridges(graph.Square)
	// fmt.Println(b1)

	// b2 := bridges.FindBridges(graph.Trapezoid)
	// fmt.Println(b2)

	b3 := bridges.FindBridges(graph.NoBridge)
	fmt.Println(b3)

	// b4 := bridges.FindBridges(graph.Bridge)
	// fmt.Println(b4)
}
