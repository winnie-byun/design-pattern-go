package factory

import "github.com/winnie-byun/design-pattern-go/common"

// Factory method creational design pattern allows
// creating objects without having to specify the exact type
// of the object that will be created.

type ShapeType int

const (
	PointType ShapeType = 1 << iota
	LineType
	CircleType
)

func NewShape(t ShapeType, args ...int) common.Shape {
	switch t {
	case LineType:
		return common.NewLine(1, 2, 3, 4)
	case CircleType:
		return common.NewCircle(10, 10, 5)
	default:
		return common.NewPoint(1, 1)
	}
}
