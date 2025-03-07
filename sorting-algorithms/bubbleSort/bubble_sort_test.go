package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{4, 2, 9, 1, 5}, []float64{1, 2, 4, 5, 9}},
		{[]float64{10, 9, 8, 7, 6}, []float64{6, 7, 8, 9, 10}},
		{[]float64{5, 5, 5, 5, 5}, []float64{5, 5, 5, 5, 5}},
		{[]float64{}, []float64{}},
	}

	for _, test := range tests {

		inputCopy := append([]float64(nil), test.input...)

		bubbleSort(inputCopy)

		fmt.Println(inputCopy)
		fmt.Println(test.expected)

		if !reflect.DeepEqual(inputCopy, test.expected) && len(inputCopy) > 0 && len(test.expected) > 0 {
			t.Errorf("Expected %v, got %v", test.expected, test.input)
		}
	}

}
