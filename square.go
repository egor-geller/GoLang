package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (x1, y1 int){
	x1 = s.start.x + int(s.a)
	y1 = s.start.y + int(s.a)
	return x1, y1
}

func (s Square) Perimeter() (p1 int)  {
	return int(s.a)*4
}

func (s Square) Area() (s1 int) {
	return int(s.a)*int(s.a)
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
