package goprogram

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial((n - 1))
}

func ReverseNumbers(n int) int {
	rev := 0
	for n > 0 {
		digits := n % 10
		rev = rev*10 + digits
		n = n / 10
	}
	return rev
}

