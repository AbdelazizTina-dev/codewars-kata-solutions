package kata

import "sort"

func FindUniq(arr []float32) float32 {
	//sorting the array, reason is the unique item will be either the first or last item :*
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if arr[len(arr)-1] != arr[len(arr)-2] {
		return arr[len(arr)-1]
	}

	return arr[0]
}
