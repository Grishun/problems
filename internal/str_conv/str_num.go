package str_conv

import (
	"math"
)

func StrToInt(str string) (output int) {
	if str[0] == '-' {
		for i := 1; i < len(str); i++ {
			output -= (int(str[i]) - 48) * int(math.Pow10(len(str)-i-1))
		}
	} else {
		for i := 0; i < len(str); i++ {
			output += (int(str[i]) - 48) * int(math.Pow10(len(str)-i-1))
		}
	}
	return output
}
