package goprogram

func FindLargestNumber(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	largest := nums[0]
	for _, num := range nums {
		if num > largest {
			largest = num
		}

	}
	return largest, true
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
