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

//func RemoveRepeating(arr []string) (output []string) {
//	i := 0
//	for j := 1; j < len(arr); j++ {
//		if i == len(arr) {
//			break
//		}
//		if arr[i] == arr[j] {
//			arr = append(arr[:j], arr[j+1:]...)
//		}
//		output = append(output, arr[i])
//		i++
//	}
//	return
//}

func RemoveRepeating(arr []string) (output []string) {
	var unique = make(map[string]struct{}, len(arr))
	for _, v := range arr {
		unique[v] = struct{}{}
	}
	for i, _ := range unique {
		output = append(output, i)
	}
	return
}

func RemoveRepeatingV2(arr []string) []string {
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
