package b_max_sum_of_contiguous_array

import (
	"fmt"
	"testing"
)

// Simple Basic Test Case
func TestFindMaxSum(t *testing.T) {
	input := []int{2, 1, 5, 1, 3, 2}
	k := 3
	expectedOutput := 9

	output := FindMaxSum(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Empty Input
func TestFindMaxSum2(t *testing.T) {
	var input []int
	k := 3
	expectedOutput := MIN_INT

	output := FindMaxSum(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// k > len(input)
func TestFindMaxSum3(t *testing.T) {
	input := []int{2, 1, 5, 1, 3, 2}
	k := 7
	expectedOutput := MIN_INT

	output := FindMaxSum(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}
