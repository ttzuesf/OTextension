package crypt

import (
	"crypto/sha256"
	"encoding/json"
	"math/big"
)

func Hash(num *big.Int) []byte {
	h := sha256.New()
	buf, _ := json.Marshal(num)
	h.Write(buf)
	return h.Sum(nil)
}
