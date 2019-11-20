package main

import "fmt"

func reverse(x []int64) []int64 {
	var l, p int = len(x), len(x) - 1
	s := make([]int64, l)
	for i := 0; i < l; i++ {
		s[p] = x[i]
		p--
	}
	return s
}

func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reverse(a[1:4]))
}
