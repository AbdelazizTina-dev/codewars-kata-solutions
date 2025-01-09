package kata

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func LongestSlideDown(pyramid [][]int) int {
	//we start from the penultimate array, add each number to the max of the two below it
	for i := len(pyramid) - 2; i >= 0; i-- {
		for j, _ := range pyramid[i] {
			pyramid[i][j] += max(pyramid[i+1][j], pyramid[i+1][j+1])
		}
	}

	//the result will be in the top
	return pyramid[0][0]
}
