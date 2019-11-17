package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (x1 Point){
	x1.x = s.start.x + int(s.a)
	x1.y = s.start.y + int(s.a)
	return
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
