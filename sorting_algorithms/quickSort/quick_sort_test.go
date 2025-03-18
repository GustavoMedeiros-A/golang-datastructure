package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 2, 9, 1, 5}, []int{1, 2, 4, 5, 9}},
		{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},
		{[]int{5, 5, 5, 5, 5}, []int{5, 5, 5, 5, 5}},
		{[]int{}, []int{}},
	}

	for _, testCase := range testCases {

		inputCopy := append([]int(nil), testCase.input...)

		quickSort(inputCopy)

		fmt.Println(inputCopy)
		fmt.Println(testCase.expected)

		if !reflect.DeepEqual(inputCopy, testCase.expected) && len(inputCopy) > 0 && len(testCase.expected) > 0 {
			t.Errorf("Expected %v, got %v", testCase.expected, inputCopy)
		}

	}
}
