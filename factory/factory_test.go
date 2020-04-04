package factory

import (
	"log"
	"testing"

	"github.com/winnie-byun/design-pattern-go/common"
)

func TestShape(t *testing.T) {
	var s common.Shape
	s = NewShape(PointType)
	log.Println(s.ToString())

	s = NewShape(CircleType)
	log.Println(s.ToString())

	s = NewShape(LineType)
	log.Println(s.ToString())
}
