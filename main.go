package main

import (
	"fmt"
	"gopractice/goprogram"
)

func main() {
	num := []int{10, 20, 69, 1}
	largest, found := goprogram.FindLargestNumber(num)
	if found {
		fmt.Println("Largest number is:", largest)
	} else {
		fmt.Println("Array is empty")
	}
	a, b := 5, 10
	goprogram.Swap(&a, &b)
	fmt.Println("After swapping:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
