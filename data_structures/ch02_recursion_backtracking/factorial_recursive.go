package rb

// calculate factorial of a positive number
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func common(n uint) uint {
	if n == 0 {
		return 1
	}

	res := uint(1)

	for i := uint(1); i <= n; i++ {
		res *= i
	}

	return res
}