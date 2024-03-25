package numbers

import "testing"

func BenchmarkFactorialV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialV1(1000)
	}
}

func BenchmarkFactorialV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialV2(1000)
	}
}
