package kata

func ProductFib(prod uint64) [3]uint64 {
	var fib_n1 uint64 = 1
	var fib_n2 uint64 = 1

	for prod > fib_n1*fib_n2 {
		fib_n := fib_n1 + fib_n2
		fib_n2 = fib_n1
		fib_n1 = fib_n

		if prod == fib_n1*fib_n2 {
			return [3]uint64{fib_n2, fib_n1, 1}
		}
	}

	return [3]uint64{fib_n2, fib_n1, 0}
}
