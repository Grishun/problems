package ProductFib

import "problems/internal/math"

type fib struct {
	num1  uint64
	num2  uint64
	isfib bool
}

func ProductFib(prod uint64) (res []uint64) {
	for i := 1; ; i++ {
		if math.Fibonacci(uint(i))*math.Fibonacci(uint(i+1)) == prod {
			res = append(res, math.Fibonacci(uint(i)), math.Fibonacci(uint(i+1)), 1)
			break
		} else if math.Fibonacci(uint(i))*math.Fibonacci(uint(i+1)) > prod {
			res = append(res, math.Fibonacci(uint(i)), math.Fibonacci(uint(i+1)), 0)
			break
		}
	}
	return
}

func ProductFibV2(prod uint64) (res fib) {
	for i := 1; ; i++ {
		if math.Fibonacci(uint(i))*math.Fibonacci(uint(i+1)) == prod {
			res.num1, res.num2, res.isfib = math.Fibonacci(uint(i)), math.Fibonacci(uint(i+1)), true
			break
		} else if math.Fibonacci(uint(i))*math.Fibonacci(uint(i+1)) > prod {
			res.num1, res.num2, res.isfib = math.Fibonacci(uint(i)), math.Fibonacci(uint(i+1)), false
			break
		}
	}
	return
}
