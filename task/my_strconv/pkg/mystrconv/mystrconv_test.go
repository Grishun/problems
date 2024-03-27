package mystrconv

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestStrToInt(t *testing.T) {

	type expected struct {
		expectedNum int
		expectedErr error
	}
	TestCases := []struct {
		input string
		exp   expected
	}{
		{
			input: "0",
			exp: expected{
				expectedNum: 0,
				expectedErr: nil,
			},
		},
		{
			input: "",
			exp: expected{
				expectedNum: 0,
				expectedErr: errors.New("can't convert an empty string"),
			},
		},
		{
			input: "-0",
			exp: expected{
				expectedNum: 0,
				expectedErr: nil,
			},
		},
		{
			input: "9223372036854775807",
			exp: expected{
				expectedNum: math.MaxInt64,
				expectedErr: nil,
			},
		},
		{
			input: "-9223372036854775808",
			exp: expected{
				expectedNum: math.MinInt64,
				expectedErr: nil,
			},
		},
	}
	for i, testCase := range TestCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			actualNum, actualErr := StrToInt(testCase.input)
			require.Equal(t, testCase.exp.expectedNum, actualNum)
			require.Equal(t, testCase.exp.expectedErr, actualErr)
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
