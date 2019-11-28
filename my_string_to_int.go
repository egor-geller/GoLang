package main

import (
	"fmt"
	"strconv"
)

func myStringToInt(s string) (int, error) {
	str, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("wrong input")
	}
	return str, nil
}
