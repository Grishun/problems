package myfor

// MyFor is function, that realize cycle, by recursion
func MyFor(i int, condition func(int) bool, fn func()) {
	if condition(i) {
		fn()
		i++
		MyFor(i, condition, fn)
	}
}
