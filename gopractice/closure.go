package gopractice

func Closure() func() int {
	y := 10
	return func() int {
		return y

	}
}
