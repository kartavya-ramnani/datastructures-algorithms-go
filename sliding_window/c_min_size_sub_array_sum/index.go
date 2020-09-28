package c_min_size_sub_array_sum

// Given an array of positive numbers and a positive number ‘S’, find the length of the smallest contiguous subarray
// whose sum is greater than or equal to ‘S’. Return 0, if no such subarray exists.
// This uses sliding window technique without fixed size of the window.
// Linear Time Complexity, Constant Space
func FindMinSize(input []int, sum int) int {
	if input == nil || len(input) == 0 {
		return 0
	}

	currentSum := 0
	windowStart := 0
	windowSize := 0
	minSize := len(input)
	for i := 1; i < len(input); i++ {
		currentSum = currentSum + input[i]
		windowSize = windowSize + 1

		// reducing the window size until this condition is met
		if currentSum >= sum {
			for currentSum >= sum {
				if minSize > windowSize {
					minSize = windowSize
				}
				currentSum = currentSum - input[windowStart]
				windowStart = windowStart + 1
				windowSize = windowSize - 1
			}
		}
	}

	return minSize
}
