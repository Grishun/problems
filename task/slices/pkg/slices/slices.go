package slices

import "errors"

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

//func RemoveByIndexV2(arr []interface{}, i uint64) (res []interface{}, err error) {
//	switch {
//	case len(arr) == 0:
//		return nil, errors.New("the array is empty")
//	case int(i) >= len(arr):
//		return nil, errors.New("i out of range")
//	default:
//		return append(arr[:i], arr[i+1:]...), nil
//	}
//}
