package sort

func Bubble(nums []int) []int {
	for i := len(nums); i > 0; i-- {
		swaps := false
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swaps = true
			}
		}
		if !swaps {
			break
		}
	}
	return nums
}
