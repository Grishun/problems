package math

import "math"

func DigitOfNum(num, n uint64) uint64 {
	return (num%uint64(math.Pow10(int(n+1))) - num%uint64(math.Pow10(int(n)))) / uint64(math.Pow10(int(n)))
}

func DigitsOfNum(num uint64) (output []uint64) {
	for i := 0; uint64(i) < Len(num); i++ {
		output = append(output, DigitOfNum(num, Len(num)-uint64(i)-1))
	}

	return
}

func DigitsOfNumV2(num uint64) (output []uint64) {
	for i := 0; uint64(i) < LenV2(num); i++ {
		output = append(output, DigitOfNum(num, LenV2(num)-uint64(i)-1))
	}

	return
}
