package bitvector

import (
	"fmt"
	"log"
	"testing"
)

func TestMatXor(t *testing.T) {
	m := 3
	n := 4
	ary1 := make([]*Bitsvect, 0)
	ary2 := make([]*Bitsvect, 0)
	for i := 0; i < n; i++ {
		ary1 = append(ary1, SampleBitsvect(m))
		ary2 = append(ary2, SampleBitsvect(m))
	}
	A := NewMatrix(ary1)
	B := NewMatrix(ary2)
	a := A.Print(A)
	fmt.Printf("A:\n%s", a)
	b := B.Print(B)
	fmt.Printf("B:\n%s", b)
	C := new(Matrix)
	C.Xor(A, B)
	fmt.Printf("C:\n%s", C.Print(C))

}

func TestMatprint(t *testing.T) {
	a := []byte{1, 2}
	b := []byte{1, 2}
	log.Println(a[0] | b[1])
}
