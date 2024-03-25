package dectobin

import (
	"fmt"
)

func DecToBin(num uint64) (bin string) {
	if num == 0 {
		return "0"
	}

	reverse := func(str string) (reversed string) {
		for i := 0; i < len(str); i++ {
			reversed += string(str[len(str)-1-i])
		}

		return
	}

	for num >= 1 {
		bin += fmt.Sprint(num % 2)
		num /= 2
	}

	return reverse(bin)
}
