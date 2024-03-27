package sequences

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10000)
	}
}

func BenchmarkFibonacciV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciV2(10000)
	}
}

func BenchmarkFibonacciV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciV3(10000)
	}
}

func BenchmarkFibonacciV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciV4(10000)
	}
}
