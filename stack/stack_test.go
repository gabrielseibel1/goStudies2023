package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int]()

	assert.NotNil(t, s)
}

func TestEnqueue(t *testing.T) {
	s := New[int]()
	i := 1

	err := s.Push(&i)
	assert.Nil(t, err)
}

func TestDequeue(t *testing.T) {
	q := New[int]()

	j, err := q.Pop()

	assert.Nil(t, j)
	assert.NotNil(t, err)
}

func TestEnqueueDequeueOnce(t *testing.T) {
	s := New[int]()
	i := 1

	_ = s.Push(&i)
	j, err := s.Pop()

	assert.Nil(t, err)
	assert.Equal(t, &i, j)
}

func TestEnqueueDequeueLots(t *testing.T) {
	s := New[int]()

	for i := 0; i < 1_000; i++ {
		v := i
		err := s.Push(&v)
		assert.Nil(t, err)
	}
	for i := 999; i >= 0; i-- {
		v, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, *v, i)
	}

	i, err := s.Pop()
	assert.Nil(t, i)
	assert.NotNil(t, err)
}
