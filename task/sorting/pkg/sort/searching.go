package sort

import (
	"errors"
)

// This func takes an array and a value,
// and returns index of the value in the array.
// If array doesn't contain the value, func returns an error
func BinSearch(arr []int, value int) (index int, err error) {
	switch {
	case len(arr) == 0:
		return 0, errors.New("the array is empty")
	case value > arr[len(arr)-1] || value < arr[0]:
		return 0, errors.New("no value in the array")
	}
	var (
		low  = 0
		high = len(arr) - 1
		mid  = (low + high) / 2
	)

	for arr[mid] != value {
		if arr[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (low + high) / 2
	}
	if low > high {
		return 0, errors.New("no value in the array")
	}

	return mid, nil
}

func LineSearch(arr []int, value int) (index int, err error) {
	index = -1

	if len(arr) == 0 {
		return 0, errors.New("the array is empty")
	}

	for i, v := range arr {
		if v == value {
			index = i
			err = nil
		}
	}

	if index == -1 {
		err = errors.New("no value in the array")
		index = 0
	}

	return
}
