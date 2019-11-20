package main

import "fmt"

func max(x []string) string {
	var l, s, p int
	l = len(x)
	s = len(x[0])
	for i := 1; i < l; i++ {
		if s < len(x[i]) {
			s = len(x[i])
			p = i
		}
	}
	return x[p]
}

func main() {
	a := []string{"one", "two", "three", "four"}
	fmt.Println(max(a))
}
