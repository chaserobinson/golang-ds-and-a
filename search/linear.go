package search


func linear(nums []int, value int) int {
	for i, num := range nums {
		if num == value {
			return i
		}
	}
	return -1
}
