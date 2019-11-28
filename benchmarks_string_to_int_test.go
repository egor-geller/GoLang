package main

import (
	"testing"
)

func Test_First(t *testing.T) {
	numStr := "1234"
	numErr := "rtt"
	num := 1234

	str1, err1 := myStringToInt(numStr)
	if str1 != num || err1 != nil {
		t.Errorf("Error")
	}

	str2, err2 := myStringToInt(numErr)
	if str2 == num || err2 == nil {
		t.Errorf("Error")
	}

	str3, err3 := myStringToInt2(numStr)
	if str3 != num || err3 != nil {
		t.Errorf("Error")
	}

	str4, err4 := myStringToInt2(numErr)
	if str4 == num || err4 == nil {
		t.Errorf("Error")
	}
}

func BenchmarkStr1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myStringToInt("1234")
	}
}

func BenchmarkStr2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		myStringToInt2("1234")
	}
}
