package eccgroup

import (
	"log"
	"math/big"
	"testing"
)

func TestECCfield(t *testing.T) {
	ecc := NewECC(256)
	log.Println("ECC Group Order:", ecc.N)
	log.Println("Set point:==================")
	P := ecc.PowG(big.NewInt(12))
	Q := ecc.PowG(big.NewInt(104))
	log.Println(P)
	log.Println(Q)
	log.Println("P+Q:========================")
	R := ecc.Mul(P, Q)
	log.Println(R)
	log.Println("P-Q:========================")
	R1 := ecc.Div(P, Q)
	log.Println(R1)
	log.Println("cmp(P,Q)=======================")
	log.Println(ecc.Cmp(R, R1))
	log.Println("Inv(Q)========================")
	Q1 := ecc.Inverse(Q)
	R2 := ecc.Mul(Q1, Q)
	log.Println(R2)
	R3 := ecc.PowG(ecc.N)
	log.Println(R3)
}
