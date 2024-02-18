package math

func Order(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j+i < len(nums); j++ {
			if nums[i] > nums[j+i] {
				nums[i], nums[j+i] = nums[j+i], nums[i]
			}
		}
	}
	return nums
}
