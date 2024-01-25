package my_cycle

func MyFor(n, k int, fn func()) {
	if n == k+1 {
		return
	}
	fn()
	n++
	MyFor(n, k, fn)
}
