package topological

import (
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func testIsTopologicallySorted(t *testing.T, name string, graph graph.Graph, sorting []int, want bool) {
	t.Run(name, func(t *testing.T) {
		if got := IsTopologicallySorted(graph, sorting); got != want {
			t.Errorf("IsTopologicallySorted() = %v, want %v", got, want)
		}
	})
}

func TestIsTopologicallySortedSquareTrue(t *testing.T) {
	testIsTopologicallySorted(t, "TestIsTopologicallySortedSquareTrue", graph.Square, []int{0, 2, 1, 3}, true)
}

func TestIsTopologicallySortedSquareFalse(t *testing.T) {
	testIsTopologicallySorted(t, "TestIsTopologicallySortedSquareFalse", graph.Square, []int{0, 1, 3, 2}, false)
}

func TestIsTopologicallySortedTrapezoidTrue(t *testing.T) {
	testIsTopologicallySorted(t, "TestIsTopologicallySortedTrapezoidTrue", graph.Trapezoid, []int{0, 1, 2, 3, 4}, true)
}

func TestIsTopologicallySortedTrapezoidFalse(t *testing.T) {
	testIsTopologicallySorted(t, "TestIsTopologicallySortedTrapezoidFalse", graph.Trapezoid, []int{4, 2, 3, 1, 0}, false)
}

func testTopSort(t *testing.T, name string, graph graph.Graph) {
	t.Run(name, func(t *testing.T) {
		if sorting := TopSort(graph); !IsTopologicallySorted(graph, sorting) {
			t.Errorf("TopSort() = %v, not topologically sorted", sorting)
		}
	})
}

func TestTopSortSquare(t *testing.T) {
	testTopSort(t, "TestTopSortSquare", graph.Square)
}

func TestTopSortTrapezoid(t *testing.T) {
	testTopSort(t, "TestTopSortTrapezoid", graph.Trapezoid)
}

func TestTopSortLong(t *testing.T) {
	testTopSort(t, "TestTopSortLong", graph.Long)
}
