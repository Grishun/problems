package main

import (
	"problems/internal/sorting"
	"testing"
)

func BenchmarkSampleV1(b *testing.B) {
	for i := 0; i < 2; i++ {
		sorting.BubbleSort([]int{1, 2, 3, 1, 4})
	}
}

func BenchmarkSampleV2(b *testing.B) {
	for i := 0; i < 2; i++ {
		sorting.SelectionSort([]int{8, 1, 7, 4, 3, 9, 2, 5, 6, 10})
	}
}
