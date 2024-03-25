package sort

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[i] > nums[i+j] {
				nums[i], nums[i+j] = nums[i+j], nums[i]
			}
		}
	}
	return nums
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
