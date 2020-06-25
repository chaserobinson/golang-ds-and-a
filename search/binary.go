package search

func binary(nums []int, value int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := int(float64(left+right) / 2)
		if nums[middle] == value {
			return middle
		}
		if value > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
