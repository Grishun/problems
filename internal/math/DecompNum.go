package math

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
