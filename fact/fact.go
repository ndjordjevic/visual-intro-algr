package fact

func fact(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res = res * i
	}

	return res
}

func factRec(n int) int {
	if n == 0 {
		return 1
	}

	return n * factRec(n-1)
}
