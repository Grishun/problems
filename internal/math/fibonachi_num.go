package math

import "math"

func Fibonacci(n uint) uint64 {
	return uint64(math.Round((math.Pow(math.Phi, float64(n)) - math.Pow(2-math.Phi, float64(n))) / math.Sqrt(5)))
}

func FibonacciV2(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	var (
		prev1  uint64
		prev2  uint64 = 1
		result uint64
	)

	for i := 2; i <= n; i++ {
		result = prev1 + prev2
		prev1 = prev2
		prev2 = result
	}

	return result
}

func FibonacciV3(n int) uint64 {
	result := make([]int, n+1)
	result[1] = 1
	for i := 2; i < n+1; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return uint64(result[n])
}
