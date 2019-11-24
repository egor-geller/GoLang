package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func max(x []string) string {
	var stringLen, p int
	if len(x) <= 0 {
		fmt.Print("Slice has no elements")
		os.Exit(1)
	}
	stringLen = len(x)
	maxCharLen := utf8.RuneCountInString(x[0])
	for i := 1; i < stringLen; i++ {
		if maxCharLen < len(x[i]) {
			maxCharLen = len(x[i])
			p = i
		}
	}
	return x[p]
}

func main() {
	a := []string{"one", "two", "three", "four"}
	//a := []string{}
	fmt.Println(max(a))
}
