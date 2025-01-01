package kata

import (
	"fmt"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	fmt.Println(array1)
	fmt.Println(array2)

	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	//remove negatives from a1
	for i, v := range array1 {
		if v < 0 {
			array1[i] = v * -1
		}
	}

	sort.Ints(array1)
	sort.Ints(array2)

	for i, v := range array1 {
		if v*v != array2[i] {
			return false
		}
	}

	return true
}
