package divide_string

// DivStr is a function, that divides string by 2 equaled by len parts
func DivStr(str string) (res []string) {
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i += 2 {
		pair := str[i : i+2]
		res = append(res, pair)
	}

	return
}
