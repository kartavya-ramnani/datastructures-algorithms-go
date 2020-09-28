package c_min_size_sub_array_sum

import (
	"fmt"
	"testing"
)

// Simple Basic Test
func TestFindMinSize1(t *testing.T) {
	input := []int{2, 1, 5, 2, 3, 2}
	sum := 7
	expectedOutput := 2 // 5 + 2

	output := FindMinSize(input, sum)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Simple Basic Test 2
func TestFindMinSize2(t *testing.T) {
	input := []int{2, 1, 5, 2, 8}
	sum := 7
	expectedOutput := 1

	output := FindMinSize(input, sum)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Empty Slice
func TestFindMinSize3(t *testing.T) {
	var input []int
	sum := 7
	expectedOutput := 0

	output := FindMinSize(input, sum)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Entire sum greater than array sum
func TestFindMinSize4(t *testing.T) {
	input := []int{2, 1, 5, 2, 8}
	sum := 19
	expectedOutput := 5 // len(input)

	output := FindMinSize(input, sum)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}
