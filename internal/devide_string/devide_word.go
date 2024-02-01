package devide_string

func Solution(str string) (res []string) {
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		pair := str[i : i+2]
		res = append(res, pair)
	}
	return
}
