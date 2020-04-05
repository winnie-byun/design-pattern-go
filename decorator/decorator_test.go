package decorator

import (
	"log"
	"testing"

	"github.com/winnie-byun/design-pattern-go/common"
)

func TestDecorator(t *testing.T) {
	b := RunWithLog(common.Build, "build")("blockchain")
	st := RunWithLog(common.Start, "start")("blockchain")
	sp := RunWithLog(common.Stop, "stop")("blockchain")
	c := RunWithLog(common.Clean, "clean")("blockchain")

	log.Println(b)
	log.Println(st)
	log.Println(sp)
	log.Println(c)
}
