package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	st := Stack{}
	st.Push("A")
	st.Push("B")
	st.Push("C")

	var expected []interface{}
	expected = append(expected, "C", "B", "A")
	assert.EqualValues(t, expected, st.sl.SliceValues())
}

func TestStackPop(t *testing.T) {
	st := Stack{}
	st.Push("Y")
	st.Push("Z")

	assert.Equal(t, st.Pop(), "Z")
	assert.Equal(t, st.Pop(), "Y")
	assert.Nil(t, st.Pop())
}
