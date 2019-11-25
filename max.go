package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func max(x []string) string {
	var sliceLen, p int
	if len(x) <= 0 {
		fmt.Print("Slice has no elements")
		os.Exit(1)
	}
	sliceLen = len(x)
	maxCharLen := utf8.RuneCountInString(x[0])
	for i := 1; i < sliceLen; i++ {

		if maxCharLen < utf8.RuneCountInString(x[i]) {
			maxCharLen = utf8.RuneCountInString(x[i])
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
