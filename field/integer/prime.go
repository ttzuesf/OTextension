package integer

import (
	"crypto/rand"
	"errors"
	"github.com/ttzuesf/goot/crypt"
	"math/big"
)

func GeneratePrimeNumber(bitsize int) *big.Int {
	A := new(big.Int)
	one := big.NewInt(1)
	two := big.NewInt(2)
	res := new(big.Int)
	for {
		pri, _ := rand.Prime(rand.Reader, bitsize) //p=2q+1, where q is a integer number
		res.Mul(pri, one)
		res.Div(res, two)
		if res.ProbablyPrime(10000) == true {
			A.Set(pri)
			break
		}
	}
	return A
}

// EncodeNumber encodes big integers using base64
func EncodeNumber(num *big.Int) string {
	buf := num.Bytes()
	return crypt.Base64Encoding(buf)
}

// DecodeNumber encodes big integers using base64
func DecodeNumber(num string) (*big.Int, error) {
	buf, err := crypt.Base64Decoding(num)
	if err != nil {
		return nil, errors.New("error integer")
	}
	return new(big.Int).SetBytes(buf), nil
}
