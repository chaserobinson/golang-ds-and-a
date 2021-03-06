package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinary(t *testing.T) {
	type test struct {
		nums  []int
		value int
		want  int
	}
	tests := []test{
		{nums: []int{0, 1, 2, 3, 4, 5}, value: 4, want: 4},
		{nums: []int{0, 5, 32, 72, 100}, value: 5, want: 1},
		{nums: []int{0, 1, 2, 3, 4, 5}, value: 7, want: -1},
		{nums: []int{}, value: 42, want: -1},
	}
	for _, tc := range tests {
		assert.Equal(t, Binary(tc.nums, tc.value), tc.want)
	}
}
