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

func Transfer(arr []int, i, j int) {
	if i < j {
		arr = append(arr[:i], append(arr[i+1:j], append([]int{arr[i]}, arr[j:]...)...)...)
	} else {
		arr = append(arr[:j], append([]int{arr[i]}, arr[j:i]...)...)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Partition(arr []int, startSort, endSort int) {
	pivot := arr[len(arr)/2]
	pivotIndex := len(arr) / 2

	for i, j := startSort, endSort; i < j; {
		if arr[i] > pivot {
			Transfer(arr, i, pivotIndex+1)
		} else if arr[j] < pivot {
			Transfer(arr, j, pivotIndex-1)
		} else {
			i, j = i+1, j-1
		}
	}

}

func quickSort(arr []int, startSort, endSort int) {

}
