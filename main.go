package main

import (
	"fmt"
	"problems/internal/my_cycle"
)

func main() {
	my_cycle.MyFor(0, 1, func() {
		fmt.Println("jopa")
	})
}
