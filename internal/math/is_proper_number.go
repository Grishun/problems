package math

import (
	"fmt"
	"math"
	"strconv"
)

func SumDigPow(a, b uint64) (output []uint64) {

	for j := a; j <= b; j++ {
		var amount uint64
		for i, v := range DigitsOfNum(j) {
			amount += uint64(math.Pow(float64(v), float64(i+1)))
		}
		if amount == j {
			output = append(output, j)
		}
	}
	return
}

func SumDigPowV2(a, b uint64) (output []uint64) {

	for j := a; j <= b; j++ {
		var amount uint64
		for i, v := range DigitsOfNumV2(j) {
			amount += uint64(math.Pow(float64(v), float64(i+1)))
		}
		if amount == j {
			output = append(output, j)
		}
	}
	return
}

// Функция, которая проверяет, является ли число суммой своих цифр в степенях и возвращает true/false.
func isSumDigPowV3(num uint64) bool {
	strNum := fmt.Sprint(num)
	var sum uint64 = 0

	for i, digit := range strNum {
		digitInt, _ := strconv.ParseUint(string(digit), 10, 32)
		sum += powInt(digitInt, uint64(i)+1)
	}

	return sum == num
}

// Функция, которая возводит число в степень.
func powInt(x, y uint64) uint64 {
	res := uint64(1)
	for y > 0 {
		if y&1 == 1 {
			res *= x
		}
		y = y >> 1
		x *= x
	}
	return res
}

// Функция, которая возвращает список чисел, удовлетворяющих условию задачи.
func SumDigPowV3(a, b uint64) []uint64 {
	var result []uint64
	var i uint64 = a
	for ; i <= b; i++ {
		if isSumDigPowV3(i) {
			result = append(result, i)
		}
	}
	return result
}
