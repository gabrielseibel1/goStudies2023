package lc703

import (
	"fmt"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	t.Run("k=3, add [163], [2], [4], [5]", func(t *testing.T) {
		cut := Constructor(3, []int{1, 6, 3})
		a := cut.Add(2)
		b := cut.Add(4)
		c := cut.Add(5)
		fmt.Println(a, b, c)
		if a != 2 {
			t.Errorf("expected a to be 2, its %d", a)
		}
		if b != 3 {
			t.Errorf("expected b to be 3, its %d", b)
		}
		if c != 4 {
			t.Errorf("expected b to be 4, its %d", c)
		}
	})

	t.Run("k=1, add [], [2], [4], [5]", func(t *testing.T) {
		cut := Constructor(1, []int{})
		a := cut.Add(-3)
		b := cut.Add(-2)
		c := cut.Add(-4)
		d := cut.Add(0)
		e := cut.Add(4)
		fmt.Println(a, b, c, d, e)
		if exp := -3; a != exp {
			t.Errorf("expected a to be %d, its %d", exp, a)
		}
		if exp := -2; b != exp {
			t.Errorf("expected b to be %d, its %d", exp, b)
		}
		if exp := -2; c != exp {
			t.Errorf("expected c to be %d, its %d", exp, c)
		}
		if exp := 0; d != exp {
			t.Errorf("expected d to be %d, its %d", exp, d)
		}
		if exp := 4; e != exp {
			t.Errorf("expected e to be %d, its %d", exp, e)
		}
	})
}
