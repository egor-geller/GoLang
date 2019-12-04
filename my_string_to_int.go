package main

import (
	"fmt"
	"strconv"
)

func myStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
