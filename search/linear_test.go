package search

import "testing"

func TestLinear(t *testing.T) {
	type test struct {
		nums  []int
		value int
		want  int
	}
	tests := []test{
		{nums: []int{0, 1, 2, 3, 4, 5}, value: 4, want: 4},
		{nums: []int{32, 5, 0, 72, 100}, value: 5, want: 1},
		{nums: []int{0, 1, 2, 3, 4, 5}, value: 7, want: -1},
		{nums: []int{}, value: 42, want: -1},
	}
	for _, tt := range tests {
		if result := linear(tt.nums, tt.value); result != tt.want {
			t.Errorf(
				"(%v, %v) -> %v [want %v]",
				tt.nums,
				tt.value,
				result,
				tt.want,
			)
		}
	}
}
