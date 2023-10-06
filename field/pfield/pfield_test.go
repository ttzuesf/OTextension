package pfield

import (
	"crypto/rand"
	"log"
	"math/big"
	"testing"
)

func TestPrimgen(t *testing.T) {
	n := 128
	prim := GeneratePrimeNumber(n)
	log.Println(string(prim.Bytes()))
}

func TestNewPField(t *testing.T) {
	field, err := NewPfield("128")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pr:", field.Pr.String())
	log.Println("g:", field.G.String())
	log.Println("e:", field.E.String())
}

func TestSavePfield(t *testing.T) {
	field, _ := NewPfield("128")
	log.Println("field", field)
	err := SavePField("field.json", field)
	if err != nil {
		log.Fatal(err)
	}
}

func TestImportField(t *testing.T) {
	field := new(Pfield)
	err := ImportField("field.json", field)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(field)
}

func TestOperationPField(t *testing.T) {
	field := new(Pfield)
	err := ImportField("field.json", field)
	if err != nil {
		log.Fatal(err)
	}
	zero := big.NewInt(0)
	a := big.NewInt(0)
	byteslen := field.Pr.BitLen() / 8
	buf := make([]byte, byteslen)
	rand.Read(buf)
	a.SetBytes(buf).Mod(a, field.Pr)
	b := big.NewInt(0)
	rand.Read(buf)
	b.SetBytes(buf).Mod(b, field.Pr)
	c := field.Mul(a, b)
	log.Println("a", a)
	log.Println("b", b)
	log.Println("c", c)
	log.Println("Power=============================")
	b.SetInt64(4)
	c = field.Pow(a, b)
	log.Println("a", a)
	log.Println("b", b)
	log.Println("c", c)
	log.Println("Inverse=======================")
	Inva := field.Inverse(a)
	c = field.Mul(Inva, a)
	log.Println("Inv:", Inva)
	log.Println("check inv:", c)
	log.Println("Div=======================")
	d := big.NewInt(-128)
	c = field.Set(d)
	if field.Sub(c, d).Cmp(zero) == 0 {
		log.Println("Set=======================")
		log.Println("set:", d)
		log.Println("check set:", c)
	}
}
