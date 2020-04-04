package common

import "fmt"

type Shape interface {
	ToString() string
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) ToString() string {
	return fmt.Sprintf("Point (%d, %d)", p.X, p.Y)
}

type Line struct {
	A Point
	B Point
}

func NewLine(x1, y1, x2, y2 int) Line {
	a := Point{x1, y1}
	b := Point{x2, y2}
	return Line{a, b}
}

func (l Line) ToString() string {
	return fmt.Sprintf("Line {%s, %s}", l.A.ToString(), l.B.ToString())
}

type Circle struct {
	P      Point
	Radius int
}

func NewCircle(x, y, radius int) Circle {
	p := Point{x, y}
	return Circle{p, radius}
}

func (c Circle) ToString() string {
	return fmt.Sprintf("Circle {%s, %d}", c.P.ToString(), c.Radius)
}
