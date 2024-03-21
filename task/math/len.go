package math

import (
	"fmt"
	"math"
)

func Len(num uint64) (res uint64) {
	if num == 0 {
		return 0
	}
	for i := 0; num/uint64(math.Pow10(i)) >= 1; i++ {
		res++
	}
	return
}

func LenV2(num uint64) uint64 {
	return uint64(len(fmt.Sprint(num)))
}
