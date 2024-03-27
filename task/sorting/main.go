package main

import (
	"fmt"
	"problems/task/sorting/pkg/sort"
)

func main() {
	fmt.Println(sort.SelectionSort([]int{5, 4, 3, 2, 1, -1}))
}
