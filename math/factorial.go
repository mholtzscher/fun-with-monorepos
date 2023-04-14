package math

func Factorial(n int) uint64 {
	var factVal uint64 = 1
	if n < 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		factVal *= uint64(i)
	}

	return factVal
}
