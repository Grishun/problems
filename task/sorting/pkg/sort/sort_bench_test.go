package sort

import "testing"

func BenchmarkSelectionSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{-1, 1, 0, 2, 3, 4, 5, 6, 7, 8})
	}
}

func BenchmarkSelectionSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{8, 7, 6, 5, 4, 3, 2, 1, 0, -1})
	}
}

func BenchmarkSelectionSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{0, 2, -1, 1, 3, 8, 5, 7, 6, 4})
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{-1, 1, 0, 2, 3, 4, 5, 6, 7, 8})
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{8, 7, 6, 5, 4, 3, 2, 1, 0, -1})
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{0, 2, -1, 1, 3, 8, 5, 7, 6, 4})
	}
}
