package e_longest_substring_with_no_repeats

import (
	"fmt"
	"testing"
)

func TestFindLongest(t *testing.T) {
	input := "ABDEFGABEF"
	expectedOutput := 6 // ABDEFG

	output := FindLongest(input)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

func TestFindLongest2(t *testing.T) {
	input := "BBBB"
	expectedOutput := 1 // B

	output := FindLongest(input)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}

func TestFindLongest3(t *testing.T) {
	input := "GEEKSFORGEEKS"
	expectedOutput := 7 // EKSFORG

	output := FindLongest(input)
	fmt.Println("Output : ", output)

	if output != expectedOutput {
		t.Error("Output and expected output do not match : ", output, expectedOutput)
	}
}
