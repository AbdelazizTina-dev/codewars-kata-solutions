package kata

import (
	"sort"
	"strconv"
)

func digToInt(digits []string) int {
	res := ""

	for _, c := range digits {
		res += c
	}

	n, _ := strconv.ParseInt(res, 10, 64)

	return int(n)
}

func NextBigger(n int) int {
	str := strconv.Itoa(n)

	// Split string into digits
	digits := []string{}
	for _, ch := range str {
		digits = append(digits, string(ch))
	}

	// find the first decreasing pair
	i := len(digits) - 1

	for i > 0 && digits[i] <= digits[i-1] {
		i--
	}

	if i == 0 {
		return -1
	}

	//find the largest on the pair and swap
	j := len(digits) - 1

	for digits[j] <= digits[i-1] {
		j--
	}

	digits[j], digits[i-1] = digits[i-1], digits[j]

	//sort the rest
	sort.Strings(digits[i:])

	return digToInt(digits)
}
