package stack

import "github.com/chaserobinson/golang-ds-and-a/list"

// Stack is a simple abstraction over a singly linked list.
// Includes Put and Take methods.
type Stack struct {
	sl list.SinglyLinked
}

// Push a new value on the stack
func (st *Stack) Push(value interface{}) interface{} {
	elem := st.sl.Unshift(value)
	return elem.Value()
}

// Pop the top value off the stack
func (st *Stack) Pop() interface{} {
	if elem := st.sl.Shift(); elem != nil {
		return elem.Value()
	}
	return nil
}
