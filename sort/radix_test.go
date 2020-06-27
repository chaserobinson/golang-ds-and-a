package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRadix(t *testing.T) {
	type test struct {
		nums     []uint
		expected []uint
	}
	tests := []test{
		{
			nums:     []uint{},
			expected: []uint{},
		},
		{
			nums:     []uint{8},
			expected: []uint{8},
		},
		{
			nums:     []uint{5, 4, 3, 2, 1, 0},
			expected: []uint{0, 1, 2, 3, 4, 5},
		},
		{
			nums:     []uint{5, 0, 342, 2, 59, 9854, 100},
			expected: []uint{0, 2, 5, 59, 100, 342, 9854},
		},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expected, Radix(tc.nums))
	}
}
