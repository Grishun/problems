package sequences

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected uint64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    10,
			expected: 55,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, Fibonacci(testCase.input))
		})
	}
}

func TestFibonacciV2(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    10,
			expected: 55,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, FibonacciV2(testCase.input))
		})
	}
}

func TestFibonacciV3(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    10,
			expected: 55,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, FibonacciV3(testCase.input))
		})
	}
}

func TestFibonacciV4(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    10,
			expected: 55,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, FibonacciV4(testCase.input))
		})
	}
}
