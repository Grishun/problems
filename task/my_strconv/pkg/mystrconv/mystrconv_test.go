package mystrconv

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestStrToInt(t *testing.T) {
	TestCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "123",
			expected: 123,
		},
		{
			input:    "-123",
			expected: -123,
		},
	}
	for i, testCase := range TestCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, StrToInt(testCase.input))
		})
	}
}

func TestIntToStr(t *testing.T) {
	TestCases := []struct {
		input    int
		expected string
	}{

		{
			input:    0,
			expected: "0",
		},
		{
			input:    -0,
			expected: "0",
		},
		{
			input:    math.MinInt64,
			expected: "-9223372036854775808",
		},
		{
			input:    math.MaxInt64,
			expected: "9223372036854775807",
		},
	}
	for i, testCase := range TestCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IntToStr(testCase.input))
		})
	}
}
