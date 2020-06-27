package queue

import "github.com/chaserobinson/golang-ds-and-a/list"

// Queue is a simple abstraction over a singly linked list.
// Includes Queue and Dequeue methods.
type Queue struct {
	sl list.SinglyLinked
}

// Enqueue a new value at the end
func (q *Queue) Enqueue(value interface{}) interface{} {
	elem := q.sl.Push(value)
	return elem.Value()
}

// Dequeue the first value
func (q *Queue) Dequeue() interface{} {
	if elem := q.sl.Shift(); elem != nil {
		return elem.Value()
	}
	return nil
}
