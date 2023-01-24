package stack

import "errors"

type Stack[T any] interface {
	Push(*T) error
	Pop() (*T, error)
}

type sliceStack[T any] []*T

func New[T any]() Stack[T] {
	s := make(sliceStack[T], 0)
	return &s
}

func (s *sliceStack[T]) Push(t *T) error {
	*s = append(*s, t)
	return nil
}

func (s *sliceStack[T]) Pop() (*T, error) {
	if len(*s) == 0 {
		return nil, errors.New("cannot Pop() from empty queue")
	}
	t := (*s)[len(*s)-1]
	(*s)[len(*s)-1] = nil // allow for item to be garbage collected
	*s = (*s)[:len(*s)-1]
	return t, nil
}
