package factory

import (
	"log"
	"testing"
)

func TestShape(t *testing.T) {
	var s Shape
	s = NewShape(PointType, 1, 2)
	log.Println(s.ToString())

	s = NewShape(CircleType, 1, 2, 3)
	log.Println(s.ToString())

	s = NewShape(LineType, 1, 2, 3, 4)
	log.Println(s.ToString())
}
