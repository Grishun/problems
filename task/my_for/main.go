package main

import (
	"fmt"
	"github.com/Grishun/problems/task/my_for/pkg/myfor"
)

func main() {
	i := 0
	myfor.MyFor(i, func(i int) bool {
		return i <= 4
	}, func() {
		fmt.Println("x")
	})
}
