package ProductFib

import (
	"problems/internal/math"
)

func ProductFib(prod int) (res []int) {
	for i := 0; math.Fibonacci(i)*math.Fibonacci(i+1) >= prod; i++ {
		res = append(res, math.Fibonacci(i), math.Fibonacci(i+1))
	}
	return res
}
