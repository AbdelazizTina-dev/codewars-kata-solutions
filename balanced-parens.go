package kata

func build(s string, n int, res *[]string, opened int, closed int) {
	if len(s) == 2*n {
		*res = append(*res, s)
		return
	}

	if opened < n {
		build(s+"(", n, res, opened+1, closed)
	}

	if closed < opened {
		build(s+")", n, res, opened, closed+1)
	}
}

func BalancedParens(n int) []string {
	res := []string{}

	build("", n, &res, 0, 0)

	return res
}
