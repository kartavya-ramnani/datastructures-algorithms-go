package d_longest_substring_with_no_more_than_k_distinct

import (
	"fmt"
	"testing"
)

// Simple Basic Test 1
func TestFindLongest1(t *testing.T) {
	input := "araaci"
	k := 2
	expectedOutput := 4 // araa

	output := FindLongest(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Simple Basic Test 2
func TestFindLongest2(t *testing.T) {
	input := "araaci"
	k := 1
	expectedOutput := 2 // aa

	output := FindLongest(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Simple Basic Test 3
func TestFindLongest3(t *testing.T) {
	input := "cbbebi"
	k := 3
	expectedOutput := 5 // cbbeb or bbebi

	output := FindLongest(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// k = 0
func TestFindLongest4(t *testing.T) {
	input := "cbbebi"
	k := 0
	expectedOutput := 0 // not possible

	output := FindLongest(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

// Empty string
func TestFindLongest5(t *testing.T) {
	input := ""
	k := 3
	expectedOutput := 0 // empty string

	output := FindLongest(input, k)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}