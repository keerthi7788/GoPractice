package gopractice

import "fmt"

func New() {
	a := new(int)
	x := 10
	a = &x
	fmt.Println(a)
	println("Value of a:", *a) // prints 0
}
