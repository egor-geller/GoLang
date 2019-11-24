package main

import (
	"fmt"
	"os"
)

func average(x []int) float64 {
	var sum, N int
	sum = 0
	N = len(x)
	if N <= 0 {
		fmt.Print("Slice has no elements")
		os.Exit(1)
	}
	for i := 0; i < N; i++ {
		sum += x[i]
	}
	return float64(sum) / float64(N)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//a := []int{}
	fmt.Println(average(a))
}
