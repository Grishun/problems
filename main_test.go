package main

import (
	"problems/internal/math"
	"testing"
)

func BenchmarkSampleV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Fibonacci(100)
	}
}

func BenchmarkSampleV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FibonacciV2(10000)
	}
}

func BenchmarkSampleV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FibonacciV3(10000)
	}
}

func BenchmarkSampleV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.FibonacciV4(10000)
	}
}
