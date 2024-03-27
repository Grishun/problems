package mystrconv

import (
	"errors"
	"math"
	"problems/task/math/pkg/numbers"
	"unicode"
)

func StrToInt(str string) (output int, err error) {
	if len(str) == 0 {
		return 0, errors.New("can't convert an empty string")
	}
	var (
		sign = 1
		i    int
	)
	if str[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < len(str); i++ {
		if unicode.IsNumber(rune(str[i])) {
			output += (int(str[i]) - 48) * int(math.Pow10(len(str)-i-1))
		} else {
			return 0, errors.New("not a number")
		}
	}

	return output * sign, nil
}

func lenOfNum(num int) (res int) {

	for num != 0 {
		num /= 10
		res++
	}

	return
}

func IntToStr(num int) string {
	var sign string
	switch {
	case num < 0:
		num = -num
		sign = "-"
	case num == 0:
		return "0"
	}
	runes := []rune{}

	for i := 1; i <= lenOfNum(num); i++ {
		digit := numbers.DigOfNum(uint64(num), uint64(lenOfNum(num)-i))
		runes = append(runes, rune('0'+digit))
	}

	return sign + string(runes)
}
