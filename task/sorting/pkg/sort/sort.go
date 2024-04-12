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

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, start, end int) int {
	pivot := arr[len(arr)-1]
	pivotIndex := len(arr) - 1
	i := start - 1

	for j := 0; j < end; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, pivotIndex)

	return i + 1
}

func QuickSort(nums []int) []int {

	pivot := partition(nums, 0, len(nums)-1)
	if pivot > 0 {
		QuickSort(nums[:pivot])
	}
	QuickSort(nums[pivot+1:])

	return nums
}
