package e_longest_substring_with_no_repeats

// Given a string, find the length of the longest substring which has no repeating characters.
func FindLongest(input string) int {
	if input == "" {
		return 0
	}

	// emulating set here by using bool
	set := make(map[string]bool)
	windowSize := 0
	maxSize := 0
	windowStart := 0

	for i := 0; i < len(input); i++ {
		for set[string(input[i])] && windowStart < len(input) {
			set[string(input[windowStart])] = false
			windowStart = windowStart + 1
			windowSize = windowSize - 1
		}

		set[string(input[i])] = true
		windowSize = windowSize + 1
		if maxSize < windowSize {
			maxSize = windowSize
		}
	}

	return maxSize
}
