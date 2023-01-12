package eqBinaryTree

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSameTrue(t *testing.T) {
	t.Run("TestSameTrue", func(t *testing.T) {
		if Same(tree.New(1), tree.New(1)) == false {
			t.Errorf("Same(tree.New(1), tree.New(1)) = false, want true")
		}
	})
}

func TestSameFalse(t *testing.T) {
	t.Run("TestSameFalse", func(t *testing.T) {
		if Same(tree.New(1), tree.New(2)) == true {
			t.Errorf("Same(tree.New(1), tree.New(2)) = true, want false")
		}
	})
}
