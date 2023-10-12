package bitvector

import (
	"fmt"
	"log"
	"testing"
)

func TestMatMul(t *testing.T) {
	a := Bitsvect{10, 1}
	b := Bitsvect{11, 12}
	log.Println("Lenght a", Length(a))
	log.Println("Length b", Length(b))
	A := new(Matrix)
	c := A.MatMul(a, b)
	log.Print(c.m, c.n)
	smat := c.MatPrint(c)
	fmt.Println(smat)
}

func TestMatprint(t *testing.T) {
	var mat []Bitsvect
	a := Bitsvect{10, 12}
	b := Bitsvect{11, 12}
	c := Bitsvect{12, 0}
	mat = append(mat, a)
	mat = append(mat, b)
	mat = append(mat, c)
	A := NewMatrix(mat)
	smat := A.MatPrint(A)
	log.Printf("\n%s", smat)
}
