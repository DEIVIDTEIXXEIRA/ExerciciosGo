package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollatz(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{1, []int{1}},
		{6, []int{6, 3, 10, 5, 16, 8, 4, 2, 1}},
		{19, []int{19, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Collatz(%d)", testCase.input), func(t *testing.T) {
			actual := Collatz(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Collatz(%d) esperado %v, obteve %v", testCase.input, testCase.expected, actual)
			}
		})
	}
}
