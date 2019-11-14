package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (e Square) End() (x1, y1 int){
	x1 = e.start.x + int(e.a)
	y1 = e.start.y + int(e.a)
	return x1, y1
}

func (p Square) Perimeter() (p1 int)  {
	return int(p.a)*4
}

func (s Square) Area() (s1 int) {
	return int(s.a)*2
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
