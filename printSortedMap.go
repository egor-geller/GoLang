package main

import "fmt"

func printSorted(x map[int]string) {
	var l int = len(x)
	e := make(map[int]string, l)
	for k := range x {
		if k > k+1 {
			e[k+1] = e[k]
		}
	}
}

func main() {
	a := map[int]string{3: "hh", 0: "tt", 1: "rr", 2: "aa", 500: "ss", 350: "ty"}
	printSorted(a)
	fmt.Println(a)
}
