package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	type test struct {
		nums []int
		want []int
	}
	tests := []test{
		{
			nums: []int{0, 1, 2, 3, 4, 5},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			nums: []int{5, 4, 3, 2, 1, 0},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			nums: []int{37, 45, 29, 8, 88, -3},
			want: []int{-3, 8, 29, 37, 45, 88},
		},
	}
	for _, tc := range tests {
		assert.Equal(t, Bubble(tc.nums), tc.want)
	}
}
