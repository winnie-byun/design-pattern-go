package factory

import "fmt"

// Factory method creational design pattern allows
// creating objects without having to specify the exact type
// of the object that will be created.

type ShapeType int

const (
	PointType ShapeType = 1 << iota
	LineType
	CircleType
)

func NewShape(t ShapeType, args ...int) Shape {
	switch t {
	case LineType:
		return newLine(args[0], args[1], args[2], args[3])
	case CircleType:
		return newCircle(args[0], args[1], args[2])
	default:
		return newPoint(args[0], args[1])
	}
}

type Shape interface {
	ToString() string
}

type Point struct {
	X int
	Y int
}

func newPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) ToString() string {
	return fmt.Sprintf("Point (%d, %d)", p.X, p.Y)
}

type Line struct {
	A Point
	B Point
}

func newLine(x1, y1, x2, y2 int) Line {
	a := Point{x1, y1}
	b := Point{x1, y1}
	return Line{a, b}
}

func (l Line) ToString() string {
	return fmt.Sprintf("Line {%s, %s}", l.A.ToString(), l.B.ToString())
}

type Circle struct {
	P      Point
	Radius int
}

func newCircle(x, y, radius int) Circle {
	p := Point{x, y}
	return Circle{p, radius}
}

func (c Circle) ToString() string {
	return fmt.Sprintf("Circle {%s, %d}", c.P.ToString(), c.Radius)
}
