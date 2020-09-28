package a_average_of_contiguous_sub_array

import (
	"fmt"
	"reflect"
	"testing"
)

// Basic Simple Positive Validation
func TestFindAverage(t *testing.T) {
	input := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	expectedOutput := []float64{2.2, 2.8, 2.4, 3.6, 2.8}

	output := FindAverage(input, k)
	fmt.Println("Output : ", output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Empty Input Testing
func TestFindAverage2(t *testing.T) {
	var input []int
	k := 5
	expectedOutput := make([]float64, 0)

	output := FindAverage(input, k)
	fmt.Println("Output : ", output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// K > len(input)
func TestFindAverage3(t *testing.T) {
	input := []int{1, 3, 2}
	k := 5
	expectedOutput := make([]float64, 0)

	output := FindAverage(input, k)
	fmt.Println("Output : ", output)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}
