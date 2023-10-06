package field

import (
	"github.com/ttzuef/ot/goot/field/eccgroup"
	"log"
	"testing"
)

func TestNewfield(t *testing.T) {
	var b Group[*eccgroup.Point]
	c := eccgroup.NewECC(224)
	b = c
	log.Println(b.Generator())
}