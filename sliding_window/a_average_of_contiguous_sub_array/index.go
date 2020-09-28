package a_average_of_contiguous_sub_array

// Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
func FindAverage(input []int, k int) []float64 {
	// boundary conditions
	if input == nil || len(input) == 0 || k == 0 || k > len(input) {
		return make([]float64, 0)
	}

	answer := make([]float64, 0)
	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + input[i]
	}
	average := float64(sum) / float64(k)
	answer = append(answer, average)

	// sliding the window
	for i := k; i < len(input); i++ {
		sum = sum + input[i] - input[i-k]
		average = float64(sum) / float64(k)
		answer = append(answer, average)
	}

	return answer
}
