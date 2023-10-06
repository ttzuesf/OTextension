// bitvector operation
package utils

import (
	"github.com/rossmerr/bitvector"
	"math/rand"
)

func Randbitvector(nlength int) *bitvector.BitVector {
	vect := bitvector.NewBitVector(nlength)
	for i := 0; i < nlength; i++ {
		c := rand.Intn(2)
		if c == 1 {
			vect.Set(i, true)
		} else {
			vect.Set(i, false)
		}
	}
	return vect
}

func RandIntvector(nlength int, rg int) []int64 {
	vect := make([]int64, nlength)
	for i := 0; i < nlength; i++ {
		c := rand.Intn(rg)
		vect[i] = int64(c)
	}
	return vect
}
