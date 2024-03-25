package productfib

import (
	"problems/task/math/pkg/sequences"
)

type fib struct {
	fib1    uint64
	fib2    uint64
	isequal bool
}

// function productFib takes an int (prod)
// and searches for two Fibonacci numbers F(n) and F(n+1) verifying
// F(n) * F(n+1) = prod. If the numbers are found, func returns them and true.
// Else func returns 2 closest numbers, verifying F(n)*F(n+1) > prod and false
func ProductFib(prod uint64) (res fib) {
	for i := 1; ; i++ {
		if sequences.Fibonacci(uint(i))*sequences.Fibonacci(uint(i+1)) == prod {
			res.fib1, res.fib2, res.isequal = sequences.Fibonacci(uint(i)), sequences.Fibonacci(uint(i+1)), true
			break
		} else if sequences.Fibonacci(uint(i))*sequences.Fibonacci(uint(i+1)) > prod {
			res.fib1, res.fib2, res.isequal = sequences.Fibonacci(uint(i)), sequences.Fibonacci(uint(i+1)), false
			break
		}
	}
	return
}
