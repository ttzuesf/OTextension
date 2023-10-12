package bitvector

import (
	"log"
	"testing"
)

func TestXor(t *testing.T) {
	x := Bitsvect{11, 2}
	y := Bitsvect{15, 1}
	c := Xor(x, y)
	log.Println(BitstoA(c))
	c = And(x, y)
	log.Println(BitstoA(c))
	c = Or(x, y)
	log.Println(BitstoA(c))
	log.Println("x:", BitstoA(x))
	log.Println("y:", BitstoA(y))
}

func TestAnd(t *testing.T) {
	A := 236
	c := 1
	k := make([]int, 0)
	for A != 0 {
		k = append(k, A&c)
		A >>= 1
	}
	log.Println(k)
	s := 0
	for i := 0; i < len(k); i++ {
		s += k[i] * (1 << i)
	}
	b := make([]uint, 2)
	b[0] = 236
	b[1] = 236
	x := b
	bk := Bools(x)
	log.Println(bk)
	y := Set(bk)
	log.Println(y)
	bs := BitstoA(x)
	log.Println(bs)
	e := AtoBits(bs)
	log.Println(e)
}

func TestRandBitsvect(t *testing.T) {
	a := RandBitsvect(128)
	log.Println(Length(a))
	str := BitstoA(a)
	log.Println(str)
}
