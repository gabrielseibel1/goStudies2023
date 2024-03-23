package queue

import "errors"

// TODO: refactor to use generics
type Queue[T any] interface {
	Enqueue(*T) error
	Dequeue() (*T, error)
}

type (
	Item              int
	sliceQueue[T any] []*T
)

func New[T any]() Queue[T] {
	q := make(sliceQueue[T], 0)
	return &q
}

func (q *sliceQueue[T]) Enqueue(t *T) error {
	*q = append(*q, t)
	return nil
}

func (q *sliceQueue[T]) Dequeue() (*T, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty, can't dequeue anything")
	}
	t := (*q)[0]
	(*q)[0] = nil // allow memory on backing array to be garbage collected
	*q = (*q)[1:]
	return t, nil
}
