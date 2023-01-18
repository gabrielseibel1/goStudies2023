package sssp

// single source shortest path

import (
	"math"
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func testShortestPath(t *testing.T, name string, graph graph.Graph, want []int) {
	t.Run(name, func(t *testing.T) {
		if got := ShortestPath(graph); !reflect.DeepEqual(got, want) {
			t.Errorf("ShortestPath() = %v, want %v", got, want)
		}
	})
}

func TestShortestPathSquare(t *testing.T) {
	testShortestPath(t, "TestShortestPathSquare", graph.SquareDG, []int{0, 1, 2, 1})
}

func TestShortestPathTrapezoid(t *testing.T) {
	testShortestPath(t, "TestShortestPathTrapezoid", graph.TrapezoidDG, []int{0, 1, 3, 0, 2})
}

func TestShortestPathLong(t *testing.T) {
	testShortestPath(t, "TestShortestPathLong", graph.LongDG, []int{0, 2, 4, 1, math.MaxInt, math.MaxInt, 4, 0, math.MaxInt, 9, 1, 9, 5})
}
