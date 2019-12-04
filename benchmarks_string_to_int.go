package main

import (
	"fmt"
	"strconv"
)

func myStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func myStringToInt2(s string) (int, error) {
	var str int
	return return fmt.Sscanf(s, "%d", &str)
}
