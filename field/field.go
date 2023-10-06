package field

import "math/big"

type Group[T any] interface {
	Set(P T) T
	Cmp(P, Q T) int
	Mul(P, Q T) T
	Div(P, Q T) T
	Inverse(P T) T
	Generator() T
	Module() *big.Int
	Pow(P T, x *big.Int) T
	PowG(x *big.Int) T
	Order() *big.Int
}
