package slices

import (
	"errors"
	"github.com/stretchr/testify/require"
	"problems/task/my_strconv/pkg/mystrconv"
	"testing"
)

func TestRemoveByIndex(t *testing.T) {

	type input struct {
		arr []int
		i   uint64
	}

	type expected struct {
		arr []int
		err error
	}
	testCases := []struct {
		inp input
		exp expected
	}{
		{
			inp: input{
				arr: []int{},
				i:   0,
			},
			exp: expected{
				arr: []int{},
				err: errors.New("the array is empty"),
			},
		},
		{
			inp: input{
				arr: []int{1},
				i:   1,
			},
			exp: expected{
				arr: []int{},
				err: errors.New("i out of range"),
			},
		},
		{
			inp: input{
				arr: []int{1, 2, 3},
				i:   0,
			},
			exp: expected{
				arr: []int{2, 3},
				err: nil,
			},
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+mystrconv.IntToStr(i), func(t *testing.T) {
			actualRes, actualErr := RemoveByIndex(testCase.inp.arr, testCase.inp.i)
			require.Equal(t, testCase.exp.arr, actualRes)
			require.Equal(t, testCase.exp.err, actualErr)
		})
	}
}

func TestRemoveByValue(t *testing.T) {

	type input struct {
		arr []int
		i   int
	}

	testCases := []struct {
		inp      input
		expected []int
	}{
		{
			inp: input{
				arr: []int{},
				i:   0,
			},
			expected: []int{},
		},
		{
			inp: input{
				arr: []int{1, 2, 3},
				i:   1,
			},
			expected: []int{2, 3},
		},
		{
			inp: input{
				arr: []int{1, 2, 3},
				i:   0,
			},
			expected: []int{1, 2, 3},
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+mystrconv.IntToStr(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveByValue(testCase.inp.arr, testCase.inp.i))
		})
	}
}

func TestRemoveRepeating(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{},
			expected: []string{},
		},
		{
			input:    []string{"apple", "banana", "apple", "orange", "banana", "grape"},
			expected: []string{"apple", "banana", "orange", "grape"},
		},
		{
			input:    []string{"apple", "banana", "orange", "grape"},
			expected: []string{"apple", "banana", "orange", "grape"},
		},
	}
	for i, testCase := range testCases {
		t.Run("case_"+mystrconv.IntToStr(i), func(t *testing.T) {
			require.Equal(t, testCase.expected, RemoveRepeating(testCase.input))
		})
	}
}
