package numbers

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecomp(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{
			input:    0,
			expected: []int{},
		},
		{
			input:    1,
			expected: []int{1},
		},
		{
			input:    10,
			expected: []int{10, 5, 2, 1},
		},
		{
			input:    -10,
			expected: []int{10, 5, 2, 1},
		},
		{
			input:    7,
			expected: []int{7, 1},
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, Decomp(testCase.input))
		})
	}
}

func TestDigSum(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    123,
			expected: 6,
		},
		{
			input:    0,
			expected: 0,
		},
		{
			input:    -123,
			expected: 6,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, DigSum(testCase.input))
		})
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    0,
			expected: false,
		},
		{
			input:    1,
			expected: false,
		},
		{
			input:    2,
			expected: true,
		},
		{
			input:    17,
			expected: true,
		},
		{
			input:    -17,
			expected: true,
		},
		{
			input:    16,
			expected: false,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, IsPrime(testCase.input))
		})
	}
}

func TestFactorialV1(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected uint64
	}{
		{
			input:    0,
			expected: 1,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    5,
			expected: 120,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, FactorialV1(testCase.input))
		})
	}
}

func TestFactorialV2(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected uint64
	}{
		{
			input:    0,
			expected: 1,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    5,
			expected: 120,
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, FactorialV2(testCase.input))
		})
	}
}
