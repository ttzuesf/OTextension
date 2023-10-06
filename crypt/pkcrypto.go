package crypt

import (
	"math/big"
)

type PKCrypto interface {
	Get_num() *big.Int
	Get_rnd_num(bitlen int64) *big.Int
	Get_fe() *big.Int
	Get_rnd_fe() *big.Int
	Get_generator() *big.Int
	Get_rnd_generator() *big.Int //
	Num_byte_size() int64
	Get_order() *big.Int
	Get_field_size() int64 //
}
