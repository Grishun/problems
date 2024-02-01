package math

func DecToBin(num uint64) (res []uint64) {
	for num == 1 {
		if num%2 != 0 {
			res = append(res, 1)
			num = (num - 1) / 2
		} else {
			res = append(res, 0)
		}
	}
	return
}
