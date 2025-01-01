package kata

import (
	"fmt"
	"math"
)

func Comp(array1 []int, array2 []int) bool {
	fmt.Println(array1)
	fmt.Println(array2)

	if array1 == nil || array2 == nil {
		return false
	}

	count1 := make(map[int]int)
	count2 := make(map[int]int)

	for _, item := range array1 {
		count1[item] += 1
	}

	for _, item := range array2 {
		sqrt := math.Sqrt(float64(item))
		if sqrt != math.Floor(sqrt) {
			return false
		}
		count2[int(sqrt)] += 1
	}

	fmt.Println(count1)
	fmt.Println(count2)

	if len(count1) != len(count2) {
		return false
	}

	for key, value1 := range count1 {
		if value2, exists := count2[key]; !exists || math.Abs(float64(value1)) != math.Abs(float64(value2)) {
			return false
		}
	}

	return true
}
