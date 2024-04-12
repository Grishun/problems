package slices

import (
	"errors"
	"sort"
)

func RemoveByIndex(arr []int, i uint64) (res []int, err error) {
	switch {
	case len(arr) == 0:
		return []int{}, errors.New("the array is empty")
	case int(i) >= len(arr):
		return []int{}, errors.New("i out of range")
	default:
		return append(arr[:i], arr[i+1:]...), nil
	}
}

func RemoveByValue(arr []int, value int) (res []int) {
	if len(arr) == 0 {
		return []int{}
	}

	for _, num := range arr {
		if num != value {
			res = append(res, num)
		}
	}

	return
}

func Shift(arr []int, shift int) []int {
	return append(arr[shift:], arr[:shift]...)
}

func RemoveRepeating(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range arr {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

func InsertionV1(arr []int, value int) (res []int) {
	
	index := sort.SearchInts(arr, value)

	return append(arr[:index], append([]int{value}, arr[index:]...)...)

}
