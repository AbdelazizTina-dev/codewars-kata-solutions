package kata

func Snail(snailMap [][]int) []int {
	if len(snailMap) == 0 {
		return []int{}
	}

	result := []int{}

	//boundaries
	top := 0
	right := len(snailMap[0]) - 1
	bottom := len(snailMap) - 1
	left := 0

	for top <= bottom && left <= right {
		//travesing the top
		for i := left; i < right; i++ {
			result = append(result, snailMap[top][i])
		}

		//traversing the right
		for i := top; i < bottom; i++ {
			result = append(result, snailMap[i][right])
		}

		//traversing the bottom
		for i := right; i >= left; i-- {
			result = append(result, snailMap[bottom][i])
		}

		//traversing the left
		for i := bottom - 1; i > top; i-- {
			result = append(result, snailMap[i][left])
		}

		top++
		right--
		bottom--
		left++
	}

	return result
}
