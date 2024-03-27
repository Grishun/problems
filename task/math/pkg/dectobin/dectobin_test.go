package dectobin

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecToBin(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected string
	}{
		{
			input:    0,
			expected: "0",
		},
		{
			input:    2,
			expected: "10",
		},
		{
			input:    10,
			expected: "1010",
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, DecToBin(testCase.input))
		})
	}
}
