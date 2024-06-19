package numbers

import (
	"math"
	"strconv"
)

// Decomp func takes a number, and returns all its dividers (only positive)
func Decomp(num int) (dividers []int) {
	if num < 0 {
		num = -num
	} else if num == 0 {
		return []int{}
	}
	for i := 0; i < num; i++ {
		if num%(num-i) == 0 {
			dividers = append(dividers, num-i)
		}
	}
	return dividers
}

// DigSum func takes a number, and returns the sum of its digits
func DigSum(num int) int {
	if num < 0 {
		num = -num
	}
	var sum int
	for i := 0; num > 0; i++ {
		sum += num % 10
		num /= 10
	}
	return sum
}

// IsPrime func checks if a number is prime
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

// DigOfNum func takes a number, and returns nth digit of it,
// counting from the last digit (last digit has index 0)
//
// Example:
// DigOfNum(123, 0) -> 3
// DigOfNum(123, 1) -> 2
// DigOfNum(123, 2) -> 1
// DigOfNum(123, 3) -> 0
func DigOfNum(num, n uint64) uint64 {
	return (num%uint64(math.Pow10(int(n+1))) - num%uint64(math.Pow10(int(n)))) / uint64(math.Pow10(int(n)))
}

func ReverseNum(num int) (res int) {
	sign := 1
	if num < 0 {
		num = -num
		sign = -1
	}
	for i := 0; i <= len(strconv.Itoa(num))-1; i++ {
		res += int(DigOfNum(uint64(num), uint64(i))) * int(math.Pow10(len(strconv.Itoa(num))-1-i))
	}
	return res * sign
}

// FactorialV1 func takes a number, and returns its factorial
func FactorialV1(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	var res uint64 = 1

	for i := 1; uint64(i) <= num; i++ {
		res *= uint64(i)
	}

	return res
}

func FactorialV2(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	return num * FactorialV2(num-1)
}

// SumOfNumPow func sums nth powers of all nums in [startRange, endRange]
func SumOfNumPow(startRange, endRange, n int) (res float64) {
	for i := startRange; i <= endRange; i++ {
		res += math.Pow(float64(i), float64(n))
	}

	return
}

// Gcd - Gratest Common Divider
func Gcd(a, b int) int {

	if b > a {
		a, b = b, a
	}

	r := a % b

	for r != 0 {
		a, b = b, r
		r = a % b
	}

	if b < 0 {
		b = -b
	}

	return b
}

func LenOfNum(num int) (res int) {
	if num < 0 {
		num = -num
	}
	for num > 0 {
		num /= 10
		res++
	}

	return
}

// Lcm - Least Common Multiple
func Lcm(a, b int) int {
	return (a * b) / (Gcd(a, b))
}

func PrimeNums(n int) (res []int) {
	if n <= 1 {
		return
	}

	isPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p < n; p++ {
		if isPrime[p] {
			for mult := p * p; mult < n; mult += p {
				isPrime[mult] = false
			}
		}
	}

	for index, value := range isPrime {
		if value {
			res = append(res, index)
		}
	}

	return
}

func DecompNum(num int) (primeDivs, divs []int) {

	primes := PrimeNums(num + 1)

	//primeDivs := make([]int, 0)

	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			primeDivs = append(divs, primes[i])
		}
	}

	for _, v := range primeDivs {
		Num := num
		for Num%v == 0 {
			Num /= v
			divs = append(divs, Num)
		}
	}

	return primeDivs, divs
}
