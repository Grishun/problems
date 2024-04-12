package main

import "fmt"

// Генерация всех возможных перестановок элементов массива
func permutations(arr []int, start int) {
	if start == len(arr)-1 {
		fmt.Println(arr)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start] // Переставляем элементы
		permutations(arr, start+1)              // Рекурсивный вызов для остальной части массива
		arr[start], arr[i] = arr[i], arr[start] // Возвращаем исходный порядок
	}
}

func main() {
	arr := []int{1, 2, 3}
	permutations(arr, 0)
}
