/*        figures.go         */
package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	x float64
}

func (s Square) area() float64 {
	return s.x * s.x
}

func (s Square) perimeter() float64 {
	return s.x * 4
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	z := math.Pi * math.Pow(c.r, 2)
	return z
}

func (c Circle) perimeter() float64 {
	z := math.Pi * 2 * c.r
	return z
}

func main() {
	var s Figure = Square{}
	var c Figure = Circle{}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}



/*        figures_test.go         */
package main

import (
	"math"
	"testing"
)

func Test_First(t *testing.T) {
	squ := &Square{}
	if squ.area() != (squ.x * squ.x) {
		t.Error("Area is not the same")
	}
	if squ.perimeter() != (squ.x * 4) {
		t.Error("Perimeter is not the same")
	}
	if squ.x < 0 {
		t.Error("No x")
	}

	cir := &Circle{}
	if cir.area() != (math.Pi * math.Pow(cir.r, 2)) {
		t.Error("Area is not the same")
	}
	if cir.perimeter() != (math.Pi * 2 * cir.r) {
		t.Error("Perimeter is not the same")
	}
	if cir.r < 0 {
		t.Error("No x")
	}
}

func TestRunMain(t *testing.T) {
	main()
}
