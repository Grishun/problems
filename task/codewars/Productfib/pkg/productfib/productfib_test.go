package productfib

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductFib(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected fib
	}{
		{
			input:    714,
			expected: fib{21, 34, true},
		},
		{
			input:    800,
			expected: fib{34, 55, false},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			ans := ProductFib(testCase.input)
			require.Equal(t, testCase.expected, ans)
		})
	}
}
