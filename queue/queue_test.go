package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	q := New[int]()

	assert.NotNil(t, q)
}

func TestEnqueue(t *testing.T) {
	q := New[int]()
	i := 1

	err := q.Enqueue(&i)
	assert.Nil(t, err)
}

func TestDequeue(t *testing.T) {
	q := New[int]()

	j, err := q.Dequeue()

	assert.Nil(t, j)
	assert.NotNil(t, err)
}

func TestEnqueueDequeueOnce(t *testing.T) {
	q := New[int]()
	i := 1

	_ = q.Enqueue(&i)
	j, err := q.Dequeue()

	assert.Nil(t, err)
	assert.Equal(t, &i, j)
}

func TestEnqueueDequeueLots(t *testing.T) {
	q := New[int]()

	for i := 0; i < 1_000; i++ {
		v := i
		err := q.Enqueue(&v)
		assert.Nil(t, err)
	}
	for i := 0; i < 1_000; i++ {
		v, err := q.Dequeue()
		assert.Nil(t, err)
		assert.Equal(t, *v, i)
	}

	i, err := q.Dequeue()
	assert.Nil(t, i)
	assert.NotNil(t, err)
}
