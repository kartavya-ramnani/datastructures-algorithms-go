package b_max_sum_of_contiguous_array

const (
	MIN_INT = -1 << 31
)


// Given an array of positive numbers and a positive number ‘k’,
// find the maximum sum of any contiguous subarray of size ‘k’.
func FindMaxSum(input []int, k int) int {
	// boundary conditions
	if input == nil || len(input) == 0 || k == 0 || k > len(input) {
		return MIN_INT
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + input[i]
	}
	maxSum := sum

	// sliding the window
	for i := k; i < len(input); i++ {
		sum = sum + input[i] - input[i-k]
		if maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum
}
