package sort

import (
	"errors"
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

func TestBinSearch(t *testing.T) {
	type input struct {
		inputArr []int
		inputVal int
	}
	type expected struct {
		expectedIndex int
		expectedErr   error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				inputArr: []int{},
				inputVal: 0,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("the array is empty"),
			},
		},
		{
			inp: input{
				inputArr: []int{1, 2, 3, 4, 5},
				inputVal: 6,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("no value in the array"),
			},
		},
		{
			inp: input{
				inputArr: []int{1, 2, 3, 4, 5},
				inputVal: -1,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("no value in the array"),
			},
		},

		{
			inp: input{
				inputArr: []int{math.MinInt64, math.MaxInt64, 0, -0},
				inputVal: math.MinInt64,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   nil,
			},
		},
		{
			inp: input{
				inputArr: []int{0, -0},
				inputVal: 0,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   nil,
			},
		},
		{
			inp: input{
				inputArr: []int{0, -0},
				inputVal: -0,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   nil,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			actualArr, actualErr := BinSearch(testCase.inp.inputArr, testCase.inp.inputVal)
			require.Equal(t, testCase.exp.expectedIndex, actualArr)
			require.Equal(t, testCase.exp.expectedErr, actualErr)

		})
	}
}

func TestLineSearch(t *testing.T) {
	type input struct {
		inputArr []int
		inputVal int
	}
	type expected struct {
		expectedIndex int
		expectedErr   error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				inputArr: []int{},
				inputVal: 0,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("the array is empty"),
			},
		},
		{
			inp: input{
				inputArr: []int{3, 5, 7, 1, 2},
				inputVal: 6,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("no value in the array"),
			},
		},
		{
			inp: input{
				inputArr: []int{1, 25, 6, 7, 3},
				inputVal: -1,
			},
			exp: expected{
				expectedIndex: 0,
				expectedErr:   errors.New("no value in the array"),
			},
		},

		{
			inp: input{
				inputArr: []int{math.MaxInt64, math.MinInt64, 0, -0},
				inputVal: math.MinInt64,
			},
			exp: expected{
				expectedIndex: 1,
				expectedErr:   nil,
			},
		},
		{
			inp: input{
				inputArr: []int{0, 2, 1, -1},
				inputVal: -1,
			},
			exp: expected{
				expectedIndex: 3,
				expectedErr:   nil,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			actualArr, actualErr := LineSearch(testCase.inp.inputArr, testCase.inp.inputVal)
			require.Equal(t, testCase.exp.expectedIndex, actualArr)
			require.Equal(t, testCase.exp.expectedErr, actualErr)

		})
	}
}
