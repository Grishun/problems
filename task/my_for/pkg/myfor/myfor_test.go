package myfor

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMyFor(t *testing.T) {
	var actualIterations int
	type cycle struct {
		i         int
		condition func(int) bool
		fn        func()
	}
	testCases := []struct {
		input    cycle
		expected int
	}{
		{
			input: cycle{
				i: 0,
				condition: func(i int) bool {
					return i <= 5
				},
				fn: func() {
					t.Logf("current = %d", actualIterations)
					actualIterations++
				},
			},
			expected: 6,
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+fmt.Sprint(i), func(t *testing.T) {
			t.SkipNow() // TODO: fix this test
			require.Equal(t, testCase.expected, actualIterations)
		})
		actualIterations = 0
	}
}
