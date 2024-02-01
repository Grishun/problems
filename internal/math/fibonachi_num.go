package math

import "math"

func Prototype(n uint) uint64 {
	return uint64(math.Round((math.Pow(math.Phi, float64(n)) - math.Pow(2-math.Phi, float64(n))) / math.Sqrt(5)))
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev1 := 0
	prev2 := 1
	result := 0

	for i := 2; i <= n; i++ {
		result = prev1 + prev2
		prev1 = prev2
		prev2 = result
	}

	return result
}

func FibbonacciV2(n int) int {
	result := make([]int, n+1)
	result[1] = 1
	for i := 2; i < n+1; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}
