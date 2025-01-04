package kata

import "fmt"

func PrimeFactors(n int) string {
	res := ""

	d := 2
	limit := n

	for d*d <= limit {
		pow := 0

		for n%d == 0 {
			pow += 1
			n = n / d
		}

		switch pow {
		case 0:
			d += 1
			continue
		case 1:
			res += fmt.Sprintf("(%v)", d)
		default:
			res += fmt.Sprintf("(%v**%v)", d, pow)
		}

		d += 1
	}

	if n > 1 {
		res += fmt.Sprintf("(%v)", n)
	}

	return res
}
