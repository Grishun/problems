package main

import (
	"fmt"
	"github.com/Grishun/problems/task/sorting/pkg/sort"
)

func main() {
	arr := []int{1, 3, 2, 1, 2, 3, 2, 1}

	sort.Partition(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
