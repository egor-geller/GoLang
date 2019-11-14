package main

import "fmt"

func factorial(i uint) uint{
	var s, x uint
	s = 1
	for x = 1; x < i ; x++{
		s *= x + 1
	}
	return s
}

func main() {
	fmt.Println(factorial(5))
}
