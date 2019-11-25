package main

import (
	"fmt"
	"os"
)

func average(x []int) (float64, error) {
	var sum, N int
	N = len(x)
	if N <= 0 {
		return 0, fmt.Errorf("slice has no elements")
	}
	for i := 0; i < N; i++ {
		sum += x[i]
	}
	return float64(sum) / float64(N), nil
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//a := []int{}
	fmt.Println(average(a))
}
