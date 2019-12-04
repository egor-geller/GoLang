package main

import (
	"testing"
)

func Test_First(t *testing.T) {
	numStrPos := "1234"
	numStrNeg := "-1234"
	numErr := "rtt"
	numPos := 1234
	numNeg := -1234

	strPos, errPos := myStringToInt(numStrPos)
	if strPos != numPos || errPos != nil {
		t.Errorf("Error")
	}

	strNeg, errNeg := myStringToInt(numStrNeg)
	if strNeg != numNeg || errNeg != nil {
		t.Errorf("Error")
	}

	str2, err2 := myStringToInt(numErr)
	if str2 == numPos || err2 == nil {
		t.Errorf("Error")
	}
}
