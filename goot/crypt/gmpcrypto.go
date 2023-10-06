package crypt

import (
	"github.com/ttzuef/ot/goot/field/pfield"
	"math/big"
)

type Gmpcrypto struct {
	Primefield *pfield.Pfield
}

func (gmp Gmpcrypto) Get_num() *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_rnd_num(bitlen int64) *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_fe() *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_rnd_fe() *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_generator() *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_rnd_generator() *big.Int {
	return nil
} //
func (gmp Gmpcrypto) Num_byte_size() int64 {
	return 0
}
func (gmp Gmpcrypto) Get_order() *big.Int {
	return nil
}
func (gmp Gmpcrypto) Get_field_size() int64 {
	return 0
}
