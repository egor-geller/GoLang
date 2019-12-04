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
	_, err := fmt.Sscanf(s, "%d", &str)
	if err != nil {
		return 0, fmt.Errorf("wrong input")
	}
	return str, nil
}
