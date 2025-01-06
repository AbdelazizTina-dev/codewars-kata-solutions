package kata

func Score(dice [5]int) int {
	count := make(map[int]int)

	for _, n := range dice {
		count[n] += 1
	}

	res := 0

	for key, value := range count {
		switch key {
		case 1:
			if value == 3 {
				res += 1000
			} else if value > 3 {
				res += 1000 + 100*(value-3)
			} else {
				res += 100 * value
			}
		case 2:
			if value >= 3 {
				res += 200
			}
		case 3:
			if value >= 3 {
				res += 300
			}
		case 4:
			if value >= 3 {
				res += 400
			}
		case 5:
			if value == 3 {
				res += 500
			} else if value > 3 {
				res += 500 + 50*(value-3)
			} else {
				res += 50 * value
			}
		case 6:
			if value >= 3 {
				res += 600
			}
		}
	}

	return res
}
