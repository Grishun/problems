package str_conv

import (
	"math"
)

func StringToNumber(str string) (output int) {
	var (
		power = 1
		i     int
	)
	if str[0] == '-' {
		power = -1
		i = 1
	}

	for ; i < len(str); i++ {
		output += (int(str[i]) - 48) * int(math.Pow10(len(str)-i-1))
	}

	return output * power
}
