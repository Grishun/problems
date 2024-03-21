package math

import "math"

func Decomp(num int) (output []int) {
	if num < 0 {
		num = -num
	}
	for i := 0; i < num; i++ {
		if num%(num-i) == 0 {
			output = append(output, num-i)
		}
	}
	return output
}

func DigSum(num int) int {
	var sum int
	for i := 0; num > 0; i++ {
		sum += num % 10
		num /= 10
	}
	return sum
}

func IsPrime(num int) bool {
	if num < 0 {
		num = -num
	}
	if num == 2 || num == 3 {
		return true
	} else if num == 1 || num == 0 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
