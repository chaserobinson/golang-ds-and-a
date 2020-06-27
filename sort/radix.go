package sort

import (
	"math"
)

// Radix sort a slice of unsigned integers and return result
func Radix(nums []uint) []uint {
	md := maxDigits(nums)

	for k := 0; k < md; k++ {
		buckets := [10][]uint{}
		for _, num := range nums {
			digit := digitAt(num, k)
			buckets[digit] = append(buckets[digit], num)
		}

		nums = []uint{}
		for _, bucket := range buckets {
			nums = append(nums, bucket...)
		}
	}
	return nums
}

// digitAt will return the digit at specified place
// Place variable represents the reverse index (e.g. 12345 place 1 == 4)
// (529, 2) -> 5
func digitAt(num uint, place int) int {
	// 10**2 == 100
	pow := math.Pow(10, float64(place))
	// 529 / 100 == 5.29
	div := float64(num) / pow
	// Floor(5.29) % 10 == 5
	mod := math.Mod(math.Floor(div), 10)
	return int(mod)
}

// digitCount returns the count of digits within given number
// (529) -> 3
func digitCount(num uint) int {
	// Log10(0) would return -Inf
	if num == 0 {
		return 1
	}
	// Log10(529) -> 2.7234556720351857
	log := math.Log10(float64(num))
	// Always positive, int() will floor log
	return int(log) + 1
}

// maxDigits returns the maximum number of digits in nums slice
// ([]uint{1, 29, 555, 29}) -> 3
func maxDigits(nums []uint) int {
	most := 0
	for _, num := range nums {
		if count := digitCount(num); count > most {
			most = count
		}
	}
	return most
}
