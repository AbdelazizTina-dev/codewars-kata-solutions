package kata

func MaximumSubarraySum(numbers []int) int {
	max := 0    // Maximum sum encountered so far
	curSum := 0 // Running sum of the current subarray

	for _, v := range numbers {
		curSum += v // Add current element to the running sum

		if curSum > max { // Update max if curSum exceeds it
			max = curSum
		}

		if curSum < 0 { // Reset curSum if it goes negative
			curSum = 0
		}
	}

	return max
}
