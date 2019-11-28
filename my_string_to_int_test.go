package main

import (
	"testing"
)

func Test_First(t *testing.T) {
	numStr := "1234"
	numErr := "rtt"
	num := 1234

	str1,err1 := myStringToInt(numStr)
	if str1 != num || err1 != nil{
		t.Errorf("Error")
	}

	str2, err2 := myStringToInt(numErr)
	if str2 == num || err2 == nil{
		t.Errorf("Error")
	}
}
