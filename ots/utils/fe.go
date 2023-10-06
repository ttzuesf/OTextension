package utils

import (
	"github.com/ttzuef/ot/goot/field/eccgroup"
	"github.com/ttzuef/ot/goot/field/pfield"
)

type Field_element interface {
	pfield.Pfield | eccgroup.ECCfield
}
