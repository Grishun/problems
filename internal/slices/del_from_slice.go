package slices

func RemoveByIndex(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func RemoveByValue(arr []int, value int) []int {
	for i, v := range arr {
		if v == value {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func Shift(arr []int, shift int) []int {
	return append(arr[shift:], arr[:shift]...)
}
