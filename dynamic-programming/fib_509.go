package dynamic_programming

func fibRecursion(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibRecursion(n-1) + fibRecursion(n-2)
}

func fibPb(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	pb := make([]int, n+1)
	pb[0], pb[1] = 0, 1
	for i := 2; i <= n; i++ {
		pb[n] = pb[n-1] + pb[n-2]
	}
	return pb[n]
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	pb0, pb1 := 0, 1
	for i := 2; i <= n; i++ {
		tmp := pb1 + pb0
		pb0 = pb1
		pb1 = tmp
	}
	return pb1
}
