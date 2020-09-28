package d_longest_substring_with_no_more_than_k_distinct

// Given a string, find the length of the longest substring in it with no more than K distinct characters
// assuming string contains lowercase characters.
func FindLongest(input string, k int) int {
	if len(input) == 0 || k == 0 {
		return 0
	}

	set := make(map[string]int)
	windowStart := 0
	windowSize := 0
	maxSize := 0

	for i := 0; i < len(input); i++ {
		set[string(input[i])] = set[string(input[i])] + 1
		windowSize++
		for len(set) > k {
			set[string(input[windowStart])] = set[string(input[windowStart])] - 1
			if set[string(input[windowStart])] == 0 {
				delete(set, string(input[windowStart]))
			}
			windowStart++
			windowSize--
		}

		if maxSize < windowSize {
			maxSize = windowSize
		}
	}

	return maxSize

}