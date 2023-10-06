package crypt

import (
	"log"
	"testing"
)

func TestGmpcrypto(t *testing.T) {
	var pk PKCrypto
	var gmp Gmpcrypto
	pk = gmp
	log.Println(pk)
}
