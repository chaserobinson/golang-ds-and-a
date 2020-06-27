package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertList(t *testing.T, sl *SinglyLinked, values ...interface{}) {
	a := assert.New(t)

	a.Equal(values, sl.SliceValues())

	length := len(values)
	a.Equal(length, sl.length)

	if length == 0 {
		a.Nil(sl.head)
		a.Nil(sl.tail)
	} else if length == 1 {
		a.Equal(sl.head, sl.tail)
	} else {
		a.NotEqual(sl.head, sl.tail)
	}
	if sl.head != nil {
		a.Nil(sl.tail.next)
	}
}

func TestSinglyLinkedPush(t *testing.T) {
	sl := NewSinglyLinked()
	sl.Push("How")
	sl.Push("are")
	sl.Push("you?")

	assertList(t, sl, "How", "are", "you?")
}

func TestSinglyLinkedPop(t *testing.T) {
	sl := NewSinglyLinked("How", "are", "you?")

	popped := sl.Pop()
	assert.Equal(t, "you?", popped.value)
	assert.Nil(t, popped.next)
	assertList(t, sl, "How", "are")

	assert.Equal(t, "are", sl.Pop().value)
	assertList(t, sl, "How")

	assert.Equal(t, "How", sl.Pop().value)
	assertList(t, sl)

	assert.Nil(t, sl.Pop())
}

func TestSinglyLinkedShift(t *testing.T) {
	sl := NewSinglyLinked("Howdy", "y'all", "!")

	unshifted := sl.Shift()
	assert.Equal(t, "Howdy", unshifted.value)
	assert.Nil(t, unshifted.next)
	assertList(t, sl, "y'all", "!")

	sl.Shift()
	assertList(t, sl, "!")

	sl.Shift()
	assertList(t, sl)

	assert.Nil(t, sl.Shift())
}

func TestSinglyLinkedUnshift(t *testing.T) {
	sl := NewSinglyLinked()

	sl.Unshift("Hey!")
	assertList(t, sl, "Hey!")

	sl.Unshift("You!")
	assertList(t, sl, "You!", "Hey!")
}

func TestSinglyLinkedGet(t *testing.T) {
	sl := NewSinglyLinked()
	assert.Nil(t, sl.Get(0))

	sl = NewSinglyLinked("The", "small", "brown", "fox")
	assert.Nil(t, sl.Get(4))
	assert.Equal(t, "brown", sl.Get(2).value)
}

func TestSinglyLinkedSet(t *testing.T) {
	sl := NewSinglyLinked()
	assert.Nil(t, sl.Set(0, "Test"))

	sl = NewSinglyLinked("The", "small", "brown", "fox")
	sl.Set(1, "big")
	assertList(t, sl, "The", "big", "brown", "fox")
}

func TestSinglyLinkedInsert(t *testing.T) {
	sl := NewSinglyLinked()
	assert.Panics(t, func() { sl.Insert(-1, "Hello") })
	assert.Panics(t, func() { sl.Insert(1, "World") })

	sl.Insert(0, "My")
	assertList(t, sl, "My")

	sl.Insert(1, "dog")
	assertList(t, sl, "My", "dog")

	sl.Insert(1, "large")
	assertList(t, sl, "My", "large", "dog")

	sl.Insert(2, "red")
	assertList(t, sl, "My", "large", "red", "dog")
}

func TestSinglyLinkedRemove(t *testing.T) {
	sl := NewSinglyLinked("My", "large", "brown", "dog")
	assert.Panics(t, func() { sl.Remove(-1) })
	assert.Panics(t, func() { sl.Remove(4) })

	removed := sl.Remove(1)
	assert.Nil(t, removed.next)
	assert.Equal(t, "large", removed.value)
	assertList(t, sl, "My", "brown", "dog")
}

func TestSinglyLinkedReverse(t *testing.T) {
	sl := NewSinglyLinked()
	assertList(t, sl.Reverse())

	sl = NewSinglyLinked(37)
	assertList(t, sl.Reverse(), 37)

	sl = NewSinglyLinked(0, 1, 2, 3, 4, 5)
	assertList(t, sl.Reverse(), 5, 4, 3, 2, 1, 0)
}
