package fibonacci

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	dp1 := 1
	dp2 := 1

	for i := 2; i < n; i++ {
		dp2, dp1 = dp1+dp2, dp2
	}

	return dp2
}
