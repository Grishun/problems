package dividestring

// DivStr is a function, that divides string to even numbers of substrings.
// If string has odd number of symbols, function adds `_` to the end of string.
//
// Example:
// "123" -> ["12", "3_"]
func DivStr(str string) (res []string) {
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}

	return
}
