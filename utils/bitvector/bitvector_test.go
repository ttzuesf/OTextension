package bitvector

import (
	"bytes"
	"log"
	"testing"
)

func TestXor(t *testing.T) {
	a := NewBitsvector("101011111111111111110011010101001010100100010100110100111000111")
	b := NewBitsvector("111011000111100001110000001111100001110000011100001110000111101")
	c := new(Bitsvect)
	c.Xor(a, b)
	c.Print(c)
	d := new(Bitsvect)
	d.And(a, b)
	d.Print(d)
}

func TestSet(t *testing.T) {
	a := new(Bitsvect)
	a.Set("101011101")
	a.Print(a)
}

func TestRandBitsvect(t *testing.T) {
	a := []byte{0, 1}
	for _, v := range a {
		log.Printf("%c\n", v)
		log.Printf("%c\n", byte('1'))
	}
}

func TestNewBitsvector(t *testing.T) {
	b := "10101010101010101010101010101001010001110001110001110001100110011000110001100011100010011001111100111"
	log.Println(len(b))
	a := NewBitsvector(b)
	c := a.String()
	if bytes.Compare([]byte(b), []byte(c)) == 0 {
		log.Printf("Correct process!")
		bits := a.Bits(a)
		log.Println(bits)
	}
}
