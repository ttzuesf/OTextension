package pfield

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"testing"
)

func TestRSAgenerate(t *testing.T) {
	n := 128
	A := new(big.Int)
	one := big.NewInt(1)
	two := big.NewInt(2)
	res := new(big.Int)
	for {
		pri, _ := rand.Prime(rand.Reader, n)
		A = pri
		res.Mul(A, one)
		res.Div(res, two)
		if res.ProbablyPrime(10000) == true {
			break
		}
	}
	log.Println(A.ProbablyPrime(100000), A.BitLen())
	a, _ := json.Marshal(A)
	b := string(a)
	log.Println("prime number:", b)
	c := []byte(b)
	B := new(big.Int)
	json.Unmarshal(c, B)
	log.Println(B.Cmp(A))
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
