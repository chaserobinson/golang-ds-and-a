package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var expected []interface{}
	expected = append(expected, 1, 2, 3)
	assert.EqualValues(t, expected, q.SliceValues())
}

func TestQueueDequeue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Nil(t, q.Dequeue())
}
