package dividestring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDivStr(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "12345",
			expected: []string{"12", "34", "5_"},
		},

		{
			input:    "12345678",
			expected: []string{"12", "34", "56", "78"},
		},

		{
			input:    "1",
			expected: []string{"1_"},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			ans := DivStr(testCase.input)

			require.Equal(t, testCase.expected, ans)
		})
	}
}
