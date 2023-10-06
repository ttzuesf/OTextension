package utils

import (
	"log"
	"testing"
)

func TestRandbitvector(t *testing.T) {
	vect := Randbitvector(10)
	log.Println(vect.String())
}
