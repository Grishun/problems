package main

import (
	"fmt"
)

func DigSum(num int) (sum int) {
	if num < 0 {
		num = -num
	}
	for i := 1; num > 0; i++ {
		sum += num % 10
		num /= 10
	}
	return
}

func Product(n int) (output []int) {
	for i := 0; i <= n; i++ {
		flag := true
		for j := 2; j <= 9; j++ {
			if DigSum(i) != DigSum(i*j) {
				flag = false
			}
		}
		if flag != false {
			output = append(output, i)
		}
	}
	return
}

func main() {
	fmt.Println(Product(100))
}
