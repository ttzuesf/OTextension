package utils

import (
	"github.com/ttzuesf/goot/field/eccgroup"
	"github.com/ttzuesf/goot/field/pfield"
)

type Field_element interface {
	pfield.Pfield | eccgroup.Curve
}
