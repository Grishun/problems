package mystrconv

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrToInt("1234")
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint("1234")
	}
}

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Atoi("1234")
	}
}

func BenchmarkIntToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToStr(1234)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(1234)
	}
}
