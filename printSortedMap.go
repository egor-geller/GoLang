package main

import (
	"fmt"
	"sort"
)

func printSorted(x map[int]string) {
	keys := make([]int, 0, len(x))
	for k := range x {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Print(x[k], " ")
	}
}

func main() {
	a := map[int]string{3: "hh", 0: "tt", 1: "rr", 2: "aa", 500: "ss", 350: "ty"}
	printSorted(a)
}
