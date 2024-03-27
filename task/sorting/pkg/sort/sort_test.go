package sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	TestCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 4, 3, 2, 1, -1},
			expected: []int{-1, 1, 2, 3, 4, 5},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{-1, 1, 2, 3, 4, 5},
			expected: []int{-1, 1, 2, 3, 4, 5},
		},

		{
			input:    []int{1},
			expected: []int{1},
		},
	}
	for i, testCase := range TestCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			ans := BubbleSort(testCase.input)
			require.Equal(t, testCase.expected, ans)
		})
	}
}

func TestSelectionSort(t *testing.T) {
	TestCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 4, 3, 2, 1, -1},
			expected: []int{-1, 1, 2, 3, 4, 5},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{-1, 1, 2, 3, 4, 5},
			expected: []int{-1, 1, 2, 3, 4, 5},
		},

		{
			input:    []int{1},
			expected: []int{1},
		},
	}
	for i, testCase := range TestCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			ans := SelectionSort(testCase.input)
			require.Equal(t, testCase.expected, ans)
		})
	}
}
